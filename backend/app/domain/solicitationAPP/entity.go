package solicitationapp

import "gorm.io/gorm"

type SolicitationApp struct {
	gorm.Model
	IdFirebase   string
	ClientId     int
	ClientName   string
	Description  string
	Address      string
	Latitude     float64
	Longitude    float64
	ProfessionId int
	CategoryId   int
	// Direção do atendimento congelada na criação: se o admin trocar o flag da
	// categoria no meio do caminho, este pedido não muda de comportamento.
	ClientTravels  bool
	ProfessionalId int
	Status         string
	ProposalValue  float64
	Distance       float64
	Rating         int
	Feedback       string
}
