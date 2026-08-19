package payment_test

import (
	"encoding/json"
	"testing"

	pkgplan "construir_mais_barato/app/domain/plan"
	pkgsubscription "construir_mais_barato/app/domain/subscription"
	pkgpaymentuc "construir_mais_barato/app/usecase/payment"
	pkgprofessionaluc "construir_mais_barato/app/usecase/professional"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// O id do pagamento no MercadoPago não pode chegar ao cliente: as rotas de
// status e de bypass são públicas, e um identificador adivinhável permitiria
// ler pagamentos alheios e sujar a auditoria.
func TestRespostaDoCheckoutNaoExpoeIdDoPagamento(t *testing.T) {
	checkout, db, _, _ := montarCheckout(t)
	criarPlanoNoBanco(t, db, pkgplan.UserTypeProfessional, 9.90, 30)

	resultado, err := checkout.Execute(pkgpaymentuc.PlanCheckoutInput{
		UserType: pkgplan.UserTypeProfessional,
		TargetID: 42,
	})
	require.NoError(t, err)

	// É o que o controller devolve na rota /publica/professional/checkout/premium.
	corpo, err := json.Marshal(&pkgprofessionaluc.CheckoutPremiumOutput{
		StatusToken:  resultado.StatusToken,
		Amount:       resultado.Amount,
		QRCode:       resultado.QRCode,
		QRCodeBase64: resultado.QRCodeBase64,
		Status:       resultado.Status,
	})
	require.NoError(t, err)

	var campos map[string]any
	require.NoError(t, json.Unmarshal(corpo, &campos))

	assert.NotContains(t, campos, "paymentId", "o id do MercadoPago não pode ir para o cliente")
	assert.NotEmpty(t, campos["statusToken"])
}

func TestStatusTokenNaoEhSerializadoNaEntidade(t *testing.T) {
	corpo, err := json.Marshal(pkgsubscription.Subscription{StatusToken: "segredo"})
	require.NoError(t, err)

	assert.NotContains(t, string(corpo), "segredo",
		"o token não deve vazar por serialização acidental da entidade")
}

func TestCadaCheckoutGeraTokenDiferente(t *testing.T) {
	checkout, db, _, _ := montarCheckout(t)
	criarPlanoNoBanco(t, db, pkgplan.UserTypeSolicitation, 9.90, 0)

	primeiro, err := checkout.Execute(pkgpaymentuc.PlanCheckoutInput{
		UserType: pkgplan.UserTypeSolicitation, TargetID: 7, ReferenceID: "sol-A",
	})
	require.NoError(t, err)

	segundo, err := checkout.Execute(pkgpaymentuc.PlanCheckoutInput{
		UserType: pkgplan.UserTypeSolicitation, TargetID: 7, ReferenceID: "sol-B",
	})
	require.NoError(t, err)

	assert.NotEmpty(t, primeiro.StatusToken)
	assert.NotEqual(t, primeiro.StatusToken, segundo.StatusToken,
		"tokens repetidos permitiriam consultar o pagamento de outra solicitação")
}

// O endpoint de bypass é público e de escrita: aceitar um identificador
// arbitrário permitiria marcar assinaturas alheias como bypassed.
func TestBypassRecusaTokenDesconhecido(t *testing.T) {
	c := montarCenario(t, mpResposta{status: "pending", amount: 9.90})

	uc := pkgpaymentuc.NewRegisterBypassUC(pkgpaymentuc.RegisterBypassUCParams{
		SubscriptionService: c.subscriptionService,
	})
	uc.Assembler = pkgpaymentuc.BypassAssembler{StatusToken: "token-inventado", UserID: 1}

	_, err := uc.Execute()
	assert.ErrorIs(t, err, pkgpaymentuc.ErrPaymentNotFound)

	uc.Assembler = pkgpaymentuc.BypassAssembler{UserID: 1}
	_, err = uc.Execute()
	assert.ErrorIs(t, err, pkgpaymentuc.ErrPaymentNotFound, "token vazio não pode criar registro")
}

func TestBypassRegistraComTokenValido(t *testing.T) {
	c := montarCenario(t, mpResposta{status: "pending", amount: 9.90})
	plano := c.criarPlano(t, pkgplan.UserTypeSolicitation, 9.90, 0)
	token := c.criarAssinaturaPendente(t, 7, plano.ID, "solicitation", 9.90)

	uc := pkgpaymentuc.NewRegisterBypassUC(pkgpaymentuc.RegisterBypassUCParams{
		SubscriptionService: c.subscriptionService,
	})
	uc.Assembler = pkgpaymentuc.BypassAssembler{
		StatusToken: token, UserID: 7, SolicitationID: "sol-1", Reason: "sem confirmação",
	}

	saida, err := uc.Execute()
	require.NoError(t, err)
	assert.True(t, saida.Registered)

	assinatura, err := c.subscriptionService.FindByStatusToken(token)
	require.NoError(t, err)
	assert.Equal(t, pkgsubscription.PaymentStatusBypassed, assinatura.Status)
	assert.Equal(t, "sol-1", assinatura.ReferenceID)
}

// Um pagamento já confirmado não é bypass: o cliente só demorou a ver.
func TestBypassNaoSobrescrevePagamentoAprovado(t *testing.T) {
	c := montarCenario(t, mpResposta{status: "pending", amount: 9.90})
	plano := c.criarPlano(t, pkgplan.UserTypeSolicitation, 9.90, 0)
	token := c.criarAssinaturaPendente(t, 7, plano.ID, "solicitation", 9.90)

	assinatura, err := c.subscriptionService.FindByStatusToken(token)
	require.NoError(t, err)
	assinatura.MarkAs(pkgsubscription.PaymentStatusApproved)
	require.NoError(t, c.subscriptionService.Update(assinatura))

	uc := pkgpaymentuc.NewRegisterBypassUC(pkgpaymentuc.RegisterBypassUCParams{
		SubscriptionService: c.subscriptionService,
	})
	uc.Assembler = pkgpaymentuc.BypassAssembler{StatusToken: token, UserID: 7}

	saida, err := uc.Execute()
	require.NoError(t, err)
	assert.False(t, saida.Registered)
	assert.Equal(t, string(pkgsubscription.PaymentStatusApproved), saida.Status)
}
