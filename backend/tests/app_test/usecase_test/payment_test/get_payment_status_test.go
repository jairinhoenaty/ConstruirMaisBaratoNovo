package payment_test

import (
	"sync/atomic"
	"testing"
	"time"

	pkgplan "construir_mais_barato/app/domain/plan"
	pkgsubscription "construir_mais_barato/app/domain/subscription"
	pkgpaymentuc "construir_mais_barato/app/usecase/payment"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// montarStatusUC devolve o caso de uso de status ligado ao mesmo cenário usado
// nos testes de webhook. ttl zero desliga o limite entre consultas.
func montarStatusUC(c *cenario, ttl time.Duration) *pkgpaymentuc.GetPaymentStatusUC {
	return pkgpaymentuc.NewGetPaymentStatusUC(pkgpaymentuc.GetPaymentStatusUCParams{
		SubscriptionService:   c.subscriptionService,
		UnlockedBudgetService: c.unlockedService,
		Processor:             c.uc,
		Throttle:              pkgpaymentuc.NewPollThrottleWithTTL(ttl),
	})
}

// É o caminho que a tela do QR Code usa: nenhum webhook chega e ainda assim o
// pagamento precisa ser reconhecido.
func TestStatusReconciliaSemWebhook(t *testing.T) {
	c := montarCenario(t, mpResposta{
		status:      "pending",
		amount:      9.90,
		externalRef: "professional:1:1",
	})

	plano := c.criarPlano(t, pkgplan.UserTypeProfessional, 9.90, 30)
	profissional := c.criarProfissional(t)
	token := c.criarAssinaturaPendente(t, profissional.ID, plano.ID, "professional", 9.90)

	statusUC := montarStatusUC(c, 0)

	// Consulta enquanto o usuário ainda não pagou.
	saida, err := statusUC.Execute(token)
	require.NoError(t, err)
	assert.False(t, saida.Approved)
	assert.Equal(t, string(pkgsubscription.PaymentStatusPending), saida.Status)

	// O usuário paga o PIX. Nenhuma notificação é entregue pelo MercadoPago.
	c.resposta.status = "approved"
	c.resposta.dateApproved = "2026-08-15T10:00:00.000-03:00"

	saida, err = statusUC.Execute(token)
	require.NoError(t, err)
	assert.True(t, saida.Approved, "o polling precisa liberar mesmo sem webhook")
	assert.NotNil(t, saida.PaidAt)

	atualizado, err := c.professionalService.FindById(profissional.ID)
	require.NoError(t, err)
	require.NotNil(t, atualizado.IsPremium)
	assert.True(t, *atualizado.IsPremium, "o premium deveria ter sido liberado pelo polling")
	assert.NotNil(t, atualizado.PremiumExpiresAt)
}

func TestStatusDePagamentoDesconhecidoRetornaNaoEncontrado(t *testing.T) {
	c := montarCenario(t, mpResposta{status: "approved", amount: 9.90})

	_, err := montarStatusUC(c, 0).Execute("token-que-nao-existe")
	assert.ErrorIs(t, err, pkgpaymentuc.ErrPaymentNotFound)
}

// O endpoint de status é público: sem limite, repetir a chamada em laço faria
// o backend martelar a API do MercadoPago.
func TestThrottleLimitaConsultasAoMercadoPago(t *testing.T) {
	var consultas int64
	c := montarCenario(t, mpResposta{
		status:      "pending",
		amount:      9.90,
		externalRef: "professional:1:1",
	})
	c.aoConsultarMP = func() { atomic.AddInt64(&consultas, 1) }

	plano := c.criarPlano(t, pkgplan.UserTypeProfessional, 9.90, 30)
	profissional := c.criarProfissional(t)
	token := c.criarAssinaturaPendente(t, profissional.ID, plano.ID, "professional", 9.90)

	statusUC := montarStatusUC(c, time.Minute)

	for i := 0; i < 5; i++ {
		_, err := statusUC.Execute(token)
		require.NoError(t, err)
	}

	assert.Equal(t, int64(1), atomic.LoadInt64(&consultas),
		"cinco consultas seguidas deveriam gerar uma única chamada ao MercadoPago")
}

func TestThrottleLiberaAposOIntervalo(t *testing.T) {
	throttle := pkgpaymentuc.NewPollThrottleWithTTL(50 * time.Millisecond)

	assert.True(t, throttle.ShouldCheck(1), "a primeira consulta nunca pode ser bloqueada")
	assert.False(t, throttle.ShouldCheck(1))

	// Pagamentos diferentes não competem entre si.
	assert.True(t, throttle.ShouldCheck(2))

	time.Sleep(60 * time.Millisecond)
	assert.True(t, throttle.ShouldCheck(1))
}
