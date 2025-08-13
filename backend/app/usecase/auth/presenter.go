package auth

type AuthenticatePresenter struct {
	Token   string        `json:"token"`
	IsLoged bool          `json:"isLoged"`
	User    UserPresenter `json:"user"`
}

type UserPresenter struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Profile     *string `json:"profile"`
	Email       string  `json:"email"`
	GoogleToken string  `json:"google_token"`
}
