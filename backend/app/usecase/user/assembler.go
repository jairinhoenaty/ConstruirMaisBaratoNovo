package user_usecase

type UserAssembler struct {
	ID          uint   `json:"oid,omitempty"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"senha"`
	Profile     string `json:"perfil"`
	GoogleToken string `json:"google_token"`
}

type LoginAssembler struct {
	Email    string `json:"email"`
	Password string `json:"senha"`
}

type ResetPasswordAssembler struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
