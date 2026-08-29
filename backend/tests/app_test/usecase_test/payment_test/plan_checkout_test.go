package payment_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	pkgplan "construir_mais_barato/app/domain/plan"
	pkgsubscription "construir_mais_barato/app/domain/subscription"
	pkgpaymentuc "construir_mais_barato/app/usecase/payment"
	"construir_mais_barato/infra/adapters/gateway-payment/mercadopago"
	pkgplaninfra "construir_mais_barato/infra/database/repositories/plan"
	pkgsubscriptioninfra "construir_mais_barato/infra/database/repositories/subscription"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// montarCheckout devolve o checkout ligado a um MercadoPago falso que gera um
// id de pagamento novo a cada chamada.
func montarCheckout(t *testing.T) (*pkgpaymentuc.PlanCheckout, *gorm.DB, pkgplan.PlanService, pkgsubscription.SubscriptionService) {
	t.Helper()

	db, err := gorm.Open(
		sqlite.Open(fmt.Sprintf("file:%s?mode=memory&cache=shared", t.Name())),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&pkgplan.Plan{}, &pkgsubscription.Subscription{}))

	t.Cleanup(func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	})

	var contador int64
	servidorMP := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := atomic.AddInt64(&contador, 1)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"id": %d,
			"status": "pending",
			"point_of_interaction": {"transaction_data": {
				"qr_code": "codigo-pix-%d",
				"qr_code_base64": "base64-%d"
			}}
		}`, id, id, id)
	}))
	t.Cleanup(servidorMP.Close)

	planService := pkgplan.NewPlanService(pkgplaninfra.NewPlanRepositoryImpl(db))
	subscriptionService := pkgsubscription.NewSubscriptionService(pkgsubscriptioninfra.NewSubscriptionRepositoryImpl(db))

	checkout := pkgpaymentuc.NewPlanCheckout(pkgpaymentuc.PlanCheckoutParams{
		MercadoPagoClient:   mercadopago.NewMPClient("token-de-teste", servidorMP.URL),
		PlanService:         planService,
		SubscriptionService: subscriptionService,
	})

	return checkout, db, planService, subscriptionService
}

func criarPlanoNoBanco(t *testing.T, db *gorm.DB, userType pkgplan.UserType, preco float64, duracaoDias int) {
	t.Helper()

	plano := pkgplan.Plan{
		UserType:     userType,
		Name:         "Plano " + string(userType),
		Price:        preco,
		Features:     "[]",
		IsActive:     true,
		DurationDays: duracaoDias,
	}
	require.NoError(t, db.Create(&plano).Error)
}

func TestCheckoutGravaAssinaturaPendente(t *testing.T) {
	checkout, db, _, subscriptionService := montarCheckout(t)
	criarPlanoNoBanco(t, db, pkgplan.UserTypeProfessional, 9.90, 30)

	resultado, err := checkout.Execute(pkgpaymentuc.PlanCheckoutInput{
		UserType: pkgplan.UserTypeProfessional,
		TargetID: 42,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, resultado.QRCode)

	// É esse registro que permite ao webhook ligar o PIX pago ao usuário.
	assinatura, err := subscriptionService.FindByPaymentID(resultado.PaymentID)
	require.NoError(t, err)
	assert.Equal(t, uint(42), assinatura.UserID)
	assert.Equal(t, "professional", assinatura.UserType)
	assert.Equal(t, 9.90, assinatura.Amount)
	assert.Equal(t, pkgsubscription.PaymentStatusPending, assinatura.Status)
	assert.Equal(t, resultado.QRCode, assinatura.QRCode)
}

func TestCheckoutPremiumReaproveitaPixPendente(t *testing.T) {
	checkout, db, _, _ := montarCheckout(t)
	criarPlanoNoBanco(t, db, pkgplan.UserTypeProfessional, 9.90, 30)

	entrada := pkgpaymentuc.PlanCheckoutInput{UserType: pkgplan.UserTypeProfessional, TargetID: 42}

	primeiro, err := checkout.Execute(entrada)
	require.NoError(t, err)

	// Clicar duas vezes não pode gerar duas cobranças.
	segundo, err := checkout.Execute(entrada)
	require.NoError(t, err)

	assert.Equal(t, primeiro.PaymentID, segundo.PaymentID)
	assert.Equal(t, primeiro.QRCode, segundo.QRCode)
}

// A taxa por solicitação é avulsa: reaproveitar o PIX entre solicitações
// diferentes faria um pagamento só liberar dois atendimentos.
func TestCheckoutSolicitacaoNaoReaproveitaEntreSolicitacoes(t *testing.T) {
	checkout, db, _, _ := montarCheckout(t)
	criarPlanoNoBanco(t, db, pkgplan.UserTypeSolicitation, 9.90, 0)

	primeiro, err := checkout.Execute(pkgpaymentuc.PlanCheckoutInput{
		UserType:    pkgplan.UserTypeSolicitation,
		TargetID:    7,
		ReferenceID: "solicitacao-A",
	})
	require.NoError(t, err)

	segundo, err := checkout.Execute(pkgpaymentuc.PlanCheckoutInput{
		UserType:    pkgplan.UserTypeSolicitation,
		TargetID:    7,
		ReferenceID: "solicitacao-B",
	})
	require.NoError(t, err)

	assert.NotEqual(t, primeiro.PaymentID, segundo.PaymentID,
		"cada solicitação precisa da sua própria cobrança")

	// Dentro da mesma solicitação, o reaproveitamento continua valendo.
	repetido, err := checkout.Execute(pkgpaymentuc.PlanCheckoutInput{
		UserType:    pkgplan.UserTypeSolicitation,
		TargetID:    7,
		ReferenceID: "solicitacao-B",
	})
	require.NoError(t, err)
	assert.Equal(t, segundo.PaymentID, repetido.PaymentID)
}

func TestCheckoutRecusaPlanoInativo(t *testing.T) {
	checkout, db, _, _ := montarCheckout(t)
	criarPlanoNoBanco(t, db, pkgplan.UserTypeStore, 19.90, 30)
	require.NoError(t, db.Model(&pkgplan.Plan{}).Where("user_type = ?", "store").Update("is_active", false).Error)

	_, err := checkout.Execute(pkgpaymentuc.PlanCheckoutInput{
		UserType: pkgplan.UserTypeStore,
		TargetID: 1,
	})
	assert.Error(t, err)
}
