package solicitationapp

import "gorm.io/gorm"

type SolicitationApp struct {
	gorm.Model
	IdFirebase     string
	ClientId       int
	ClientName     string
	Description    string
	Address        string
	Latitude       float64
	Longitude      float64
	ProfessionId   int
	ProfessionalId int
	Status         string
	ProposalValue  float64
	Distance       float64
}
