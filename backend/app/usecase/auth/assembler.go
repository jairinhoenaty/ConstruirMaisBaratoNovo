package auth

type LoginAssembler struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ValidaLoginAssembler struct {
	Token string `json:"token"`
	//Password string `json:"password"`
}

type ExchangeCodeAssembler struct {
	UserId uint `json:"userId"`
}
type RedeemCodeAssembler struct {
	Code string `json:"code"`
}
