package chat_usecase

type ChatAssembler struct {
	ID             uint   `json:"id,omitempty"`
	Message        string `json:"message"`
	ProfessionalID *uint  `json:"professional_id"`
	ClientID       *uint  `json:"client_id"`
	Origem         string `json:"origem"`
}

type FindWithPaginationUserAssembler struct {
	Limit          int  `json:"limit"`
	Offset         int  `json:"offset"`
	ProfessionalID uint `json:"professional_id"`
	ClientID       uint `json:"client_id"`
}

type FindWithPaginationChatAssembler struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
