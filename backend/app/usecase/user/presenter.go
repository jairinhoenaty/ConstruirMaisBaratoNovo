package user_usecase

type UserPresenter struct {
	ID          uint   `json:"oid,omitempty"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"senha"`
	Profile     string `json:"perfil"`
	GoogleToken string `json:"google_token"`
}
