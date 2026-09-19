package accountdeletion_usecase

type CreateAccountDeletionAssembler struct {
	UserID    *uint  `json:"userId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Reason    string `json:"reason"`
}
