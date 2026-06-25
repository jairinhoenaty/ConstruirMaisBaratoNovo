package solicitation

type SaveSolicitationAssembler struct {
	IdFirebase     string  `json:"idFirebase"`
	ClientId       int     `json:"clientId"`
	ClientName     string  `json:"clientName"`
	Description    string  `json:"description"`
	Address        string  `json:"address"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	ProfessionId   int     `json:"professionId"`
	ProfessionalId int     `json:"idProfissional"`
	Status         string  `json:"status"`
	Distance       float64 `json:"distance"`
	ProposalValue  float64 `json:"proposalValue"`
}

type UpdateFeedbackAssembler struct {
	IdFirebase string `json:"idFirebase"`
	Rating     int    `json:"rating"`
	Feedback   string `json:"feedback"`
}
