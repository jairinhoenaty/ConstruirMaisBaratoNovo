package payment_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	pkgplan "construir_mais_barato/app/domain/plan"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgstore "construir_mais_barato/app/domain/store"
	pkgsubscription "construir_mais_barato/app/domain/subscription"
	pkgunlockedbudget "construir_mais_barato/app/domain/unlocked-budget"
	pkgpaymentuc "construir_mais_barato/app/usecase/payment"
	"construir_mais_barato/infra/adapters/gateway-payment/mercadopago"
	pkgplaninfra "construir_mais_barato/infra/database/repositories/plan"
	pkgprofessionalinfra "construir_mais_barato/infra/database/repositories/professional"
	pkgstoreinfra "construir_mais_barato/infra/database/repositories/store"
	pkgsubscriptioninfra "construir_mais_barato/infra/database/repositories/subscription"
	pkgunlockedbudgetinfra "construir_mais_barato/infra/database/repositories/unlocked-budget"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const paymentID int64 = 987654

type cenario struct {
	db                  *gorm.DB
	uc                  *pkgpaymentuc.ProcessPaymentNotificationUC
	resposta            *mpResposta // o teste pode ajustar a resposta depois de criar os registros
	aoConsultarMP       func()      // chamado a cada requisição recebida pelo MercadoPago falso
	subscriptionService pkgsubscription.SubscriptionService
	professionalService pkgprofessional.ProfessionalService
	storeService        pkgstore.StoreService
	planService         pkgplan.PlanService
	unlockedService     pkgunlockedbudget.UnlockedBudgetService
}

// mpResposta descreve o pagamento que o MercadoPago devolverá no teste.
type mpResposta struct {
	status       string
	amount       float64
	externalRef  string
	dateApproved string
}

func montarCenario(t *testing.T, resposta mpResposta) *cenario {
	t.Helper()

	db, err := gorm.Open(
		sqlite.Open(fmt.Sprintf("file:%s?mode=memory&cache=shared", t.Name())),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)
	require.NoError(t, err)

	require.NoError(t, db.AutoMigrate(
		&pkgplan.Plan{},
		&pkgsubscription.Subscription{},
		&pkgprofessional.Professional{},
		&pkgstore.Store{},
		&pkgunlockedbudget.UnlockedBudget{},
	))

	t.Cleanup(func() {
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
		}
	})

	atual := resposta
	c := &cenario{}
	servidorMP := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if c.aoConsultarMP != nil {
			c.aoConsultarMP()
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"id": %d,
			"status": %q,
			"status_detail": "accredited",
			"transaction_amount": %v,
			"description": "Premium",
			"payment_method_id": "pix",
			"external_reference": %q,
			"date_created": "2026-08-14T10:00:00.000-03:00",
			"date_approved": %q,
			"date_last_updated": "2026-08-14T10:05:00.000-03:00"
		}`, paymentID, atual.status, atual.amount, atual.externalRef, atual.dateApproved)
	}))
	t.Cleanup(servidorMP.Close)

	c.db = db
	c.resposta = &atual
	c.subscriptionService = pkgsubscription.NewSubscriptionService(pkgsubscriptioninfra.NewSubscriptionRepositoryImpl(db))
	c.professionalService = pkgprofessional.NewProfessionalService(pkgprofessionalinfra.NewProfessionalRepositoryImpl(db))
	c.storeService = pkgstore.NewStoreService(pkgstoreinfra.NewStoreRepositoryImpl(db))
	c.planService = pkgplan.NewPlanService(pkgplaninfra.NewPlanRepositoryImpl(db))
	c.unlockedService = pkgunlockedbudget.NewUnlockedBudgetService(pkgunlockedbudgetinfra.NewUnlockedBudgetRepositoryImpl(db))

	c.uc = pkgpaymentuc.NewProcessPaymentNotificationUC(pkgpaymentuc.ProcessPaymentNotificationUCParams{
		MercadoPagoClient:     mercadopago.NewMPClient("token-de-teste", servidorMP.URL),
		SubscriptionService:   c.subscriptionService,
		UnlockedBudgetService: c.unlockedService,
		ProfessionalService:   c.professionalService,
		StoreService:          c.storeService,
		PlanService:           c.planService,
	})

	return c
}

func (c *cenario) criarPlano(t *testing.T, userType pkgplan.UserType, preco float64, duracaoDias int) *pkgplan.Plan {
	t.Helper()

	plano := pkgplan.Plan{
		UserType:     userType,
		Name:         "Plano " + string(userType),
		Price:        preco,
		Features:     "[]",
		IsActive:     true,
		DurationDays: duracaoDias,
	}
	require.NoError(t, c.db.Create(&plano).Error)

	return &plano
}

func (c *cenario) criarProfissional(t *testing.T) *pkgprofessional.Professional {
	t.Helper()

	naoPremium := false
	profissional := pkgprofessional.Professional{
		Name:      "Profissional Teste",
		Email:     "profissional@teste.com",
		IsPremium: &naoPremium,
	}
	require.NoError(t, c.db.Create(&profissional).Error)

	return &profissional
}

// criarAssinaturaPendente devolve o token opaco usado para consultar o status.
func (c *cenario) criarAssinaturaPendente(t *testing.T, userID, planID uint, userType string, valor float64) string {
	t.Helper()

	statusToken := fmt.Sprintf("token-%s", t.Name())
	_, err := c.subscriptionService.Save(pkgsubscription.Subscription{
		StatusToken:       statusToken,
		UserID:            userID,
		PlanID:            planID,
		UserType:          userType,
		PaymentID:         paymentID,
		ExternalReference: fmt.Sprintf("%s:%d:1", userType, userID),
		Amount:            valor,
		Status:            pkgsubscription.PaymentStatusPending,
		IdempotencyKey:    fmt.Sprintf("chave-%s", t.Name()),
	})
	require.NoError(t, err)

	return statusToken
}

func TestPagamentoAprovadoLiberaPremiumDoProfissional(t *testing.T) {
	c := montarCenario(t, mpResposta{
		status:       "approved",
		amount:       9.90,
		externalRef:  "professional:1:1",
		dateApproved: "2026-08-14T10:05:00.000-03:00",
	})

	plano := c.criarPlano(t, pkgplan.UserTypeProfessional, 9.90, 30)
	profissional := c.criarProfissional(t)
	c.criarAssinaturaPendente(t, profissional.ID, plano.ID, "professional", 9.90)

	resultado, err := c.uc.Execute(paymentID)
	require.NoError(t, err)

	assert.True(t, resultado.Handled)
	assert.True(t, resultado.Activated)
	assert.Equal(t, pkgsubscription.PaymentStatusApproved, resultado.Status)

	atualizado, err := c.professionalService.FindById(profissional.ID)
	require.NoError(t, err)
	require.NotNil(t, atualizado.IsPremium)
	assert.True(t, *atualizado.IsPremium, "o profissional deveria ter virado premium")

	require.NotNil(t, atualizado.PremiumExpiresAt, "a vigência deveria ter sido preenchida")
	esperado := time.Date(2026, 8, 14, 10, 5, 0, 0, time.FixedZone("", -3*60*60)).Add(30 * 24 * time.Hour)
	assert.WithinDuration(t, esperado, *atualizado.PremiumExpiresAt, time.Minute)

	assinatura, err := c.subscriptionService.FindByPaymentID(paymentID)
	require.NoError(t, err)
	assert.Equal(t, pkgsubscription.PaymentStatusApproved, assinatura.Status)
	assert.NotNil(t, assinatura.PaidAt)
}

func TestNotificacaoRepetidaNaoReprocessa(t *testing.T) {
	c := montarCenario(t, mpResposta{
		status:       "approved",
		amount:       9.90,
		externalRef:  "professional:1:1",
		dateApproved: "2026-08-14T10:05:00.000-03:00",
	})

	plano := c.criarPlano(t, pkgplan.UserTypeProfessional, 9.90, 30)
	profissional := c.criarProfissional(t)
	c.criarAssinaturaPendente(t, profissional.ID, plano.ID, "professional", 9.90)

	primeiro, err := c.uc.Execute(paymentID)
	require.NoError(t, err)
	assert.True(t, primeiro.Activated)

	// O MercadoPago reenvia a mesma notificação várias vezes.
	segundo, err := c.uc.Execute(paymentID)
	require.NoError(t, err)
	assert.True(t, segundo.Handled)
	assert.False(t, segundo.Activated, "a segunda notificação não deveria reprocessar")
	assert.Equal(t, "status inalterado", segundo.Description)
}

func TestValorDivergenteNaoLiberaPremium(t *testing.T) {
	// O pagador alterou o valor do PIX: cobramos 9,90 e ele pagou 0,01.
	c := montarCenario(t, mpResposta{
		status:       "approved",
		amount:       0.01,
		externalRef:  "professional:1:1",
		dateApproved: "2026-08-14T10:05:00.000-03:00",
	})

	plano := c.criarPlano(t, pkgplan.UserTypeProfessional, 9.90, 30)
	profissional := c.criarProfissional(t)
	c.criarAssinaturaPendente(t, profissional.ID, plano.ID, "professional", 9.90)

	resultado, err := c.uc.Execute(paymentID)
	require.NoError(t, err)

	assert.False(t, resultado.Activated)
	assert.Equal(t, pkgsubscription.PaymentStatusFailed, resultado.Status)

	atualizado, err := c.professionalService.FindById(profissional.ID)
	require.NoError(t, err)
	require.NotNil(t, atualizado.IsPremium)
	assert.False(t, *atualizado.IsPremium, "valor divergente não pode liberar premium")
}

func TestEstornoRevogaPremium(t *testing.T) {
	c := montarCenario(t, mpResposta{
		status:      "refunded",
		amount:      9.90,
		externalRef: "professional:1:1",
	})

	plano := c.criarPlano(t, pkgplan.UserTypeProfessional, 9.90, 30)
	profissional := c.criarProfissional(t)
	c.criarAssinaturaPendente(t, profissional.ID, plano.ID, "professional", 9.90)

	// O premium já estava liberado quando veio o estorno.
	vigencia := time.Now().Add(30 * 24 * time.Hour)
	require.NoError(t, c.professionalService.SetPremium(profissional.ID, true, &vigencia))

	resultado, err := c.uc.Execute(paymentID)
	require.NoError(t, err)
	assert.Equal(t, pkgsubscription.PaymentStatusRefunded, resultado.Status)

	atualizado, err := c.professionalService.FindById(profissional.ID)
	require.NoError(t, err)
	require.NotNil(t, atualizado.IsPremium)
	assert.False(t, *atualizado.IsPremium, "o estorno deveria revogar o premium")
	assert.Nil(t, atualizado.PremiumExpiresAt)
}

func TestPagamentoPendenteNaoLiberaNada(t *testing.T) {
	c := montarCenario(t, mpResposta{
		status:      "pending",
		amount:      9.90,
		externalRef: "professional:1:1",
	})

	plano := c.criarPlano(t, pkgplan.UserTypeProfessional, 9.90, 30)
	profissional := c.criarProfissional(t)
	c.criarAssinaturaPendente(t, profissional.ID, plano.ID, "professional", 9.90)

	resultado, err := c.uc.Execute(paymentID)
	require.NoError(t, err)
	assert.False(t, resultado.Activated)

	atualizado, err := c.professionalService.FindById(profissional.ID)
	require.NoError(t, err)
	require.NotNil(t, atualizado.IsPremium)
	assert.False(t, *atualizado.IsPremium)
}

// Pagamentos criados antes deste recurso não têm linha em subscriptions: o
// vínculo precisa ser reconstruído pelo external_reference.
func TestPagamentoLegadoEhRecuperadoPeloExternalReference(t *testing.T) {
	c := montarCenario(t, mpResposta{
		status:       "approved",
		amount:       9.90,
		dateApproved: "2026-08-14T10:05:00.000-03:00",
	})

	c.criarPlano(t, pkgplan.UserTypeProfessional, 9.90, 30)
	profissional := c.criarProfissional(t)

	// Nenhuma assinatura gravada: só o external_reference aponta para o dono.
	c.resposta.externalRef = fmt.Sprintf("professional:%d:1699999999", profissional.ID)

	resultado, err := c.uc.Execute(paymentID)
	require.NoError(t, err)
	assert.True(t, resultado.Activated, "o pagamento legado deveria liberar o premium")

	atualizado, err := c.professionalService.FindById(profissional.ID)
	require.NoError(t, err)
	require.NotNil(t, atualizado.IsPremium)
	assert.True(t, *atualizado.IsPremium)

	// O registro que faltava deve ter sido criado para auditoria.
	assinatura, err := c.subscriptionService.FindByPaymentID(paymentID)
	require.NoError(t, err)
	assert.Equal(t, pkgsubscription.PaymentStatusApproved, assinatura.Status)
	assert.Equal(t, profissional.ID, assinatura.UserID)
}

// Notificações do MercadoPago podem chegar fora de ordem.
func TestNotificacaoPendenteAtrasadaNaoDesfazAprovacao(t *testing.T) {
	c := montarCenario(t, mpResposta{
		status:       "approved",
		amount:       9.90,
		externalRef:  "professional:1:1",
		dateApproved: "2026-08-14T10:05:00.000-03:00",
	})

	plano := c.criarPlano(t, pkgplan.UserTypeProfessional, 9.90, 30)
	profissional := c.criarProfissional(t)
	c.criarAssinaturaPendente(t, profissional.ID, plano.ID, "professional", 9.90)

	aprovado, err := c.uc.Execute(paymentID)
	require.NoError(t, err)
	require.True(t, aprovado.Activated)

	// Chega, atrasada, a notificação anterior com status "pending".
	c.resposta.status = "pending"

	resultado, err := c.uc.Execute(paymentID)
	require.NoError(t, err)
	assert.Equal(t, pkgsubscription.PaymentStatusApproved, resultado.Status)

	atualizado, err := c.professionalService.FindById(profissional.ID)
	require.NoError(t, err)
	assert.True(t, *atualizado.IsPremium, "o premium não pode cair por notificação atrasada")

	assinatura, err := c.subscriptionService.FindByPaymentID(paymentID)
	require.NoError(t, err)
	assert.Equal(t, pkgsubscription.PaymentStatusApproved, assinatura.Status)
}

func TestPagamentoDesconhecidoNaoQuebra(t *testing.T) {
	c := montarCenario(t, mpResposta{
		status:      "approved",
		amount:      50.00,
		externalRef: "integracao-de-terceiros",
	})

	resultado, err := c.uc.Execute(paymentID)
	require.NoError(t, err, "um pagamento de outra integração não pode derrubar o webhook")
	assert.False(t, resultado.Handled)
	assert.False(t, resultado.Activated)
}

func TestExpiracaoNaoRebaixaPremiumManual(t *testing.T) {
	c := montarCenario(t, mpResposta{status: "approved", amount: 9.90})

	// Premium com prazo vencido: deve cair.
	vencido := time.Now().Add(-24 * time.Hour)
	comPrazo := c.criarProfissional(t)
	require.NoError(t, c.professionalService.SetPremium(comPrazo.ID, true, &vencido))

	// Premium ativado na mão por um administrador, sem prazo: deve permanecer.
	semPrazo := pkgprofessional.Professional{Name: "Manual", Email: "manual@teste.com"}
	require.NoError(t, c.db.Create(&semPrazo).Error)
	require.NoError(t, c.professionalService.SetPremium(semPrazo.ID, true, nil))

	uc := pkgpaymentuc.NewExpirePremiumsUC(pkgpaymentuc.ExpirePremiumsUCParams{
		ProfessionalService: c.professionalService,
		StoreService:        c.storeService,
		SubscriptionService: c.subscriptionService,
	})

	resultado, err := uc.Execute()
	require.NoError(t, err)
	assert.Equal(t, int64(1), resultado.Professionals)

	rebaixado, err := c.professionalService.FindById(comPrazo.ID)
	require.NoError(t, err)
	assert.False(t, *rebaixado.IsPremium, "premium com prazo vencido deveria cair")

	mantido, err := c.professionalService.FindById(semPrazo.ID)
	require.NoError(t, err)
	assert.True(t, *mantido.IsPremium, "premium manual do admin não pode ser rebaixado")
}
