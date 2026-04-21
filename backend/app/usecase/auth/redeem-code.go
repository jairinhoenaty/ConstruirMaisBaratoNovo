package auth

import (
	pkgexchangecode "construir_mais_barato/app/domain/exchange-codes"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkguser "construir_mais_barato/app/domain/user"
	"errors"
)

type RedeemCodeUC struct {
	UserService         pkguser.UserService
	ProfessionalService pkgprofessional.ProfessionalService
	ExchangeCodeService pkgexchangecode.ExchangeCodeService
	Assembler           *RedeemCodeAssembler
}

type RedeemCodeUCParams struct {
	UserService         pkguser.UserService
	ProfessionalService pkgprofessional.ProfessionalService
	ExchangeCodeService pkgexchangecode.ExchangeCodeService
}

func NewRedeemCode(params RedeemCodeUCParams) RedeemCodeUC {
	return RedeemCodeUC{
		UserService:         params.UserService,
		ProfessionalService: params.ProfessionalService,
		ExchangeCodeService: params.ExchangeCodeService,
	}
}

func (uc *RedeemCodeUC) Execute() (*AuthenticatePresenter, error) {

	if uc.Assembler == nil && uc.Assembler.Code != "" {
		return nil, errors.New("Invalid data")
	}

	presenter := AuthenticatePresenter{}

	validCode, _ := uc.ExchangeCodeService.Redeem(uc.Assembler.Code)

	if validCode != nil && validCode.Code != "" {
		token, err := GenerateToken(UserPresenter{
			ID:   validCode.UserID,
			Name: validCode.User.Name,
		})

		if err == nil {
			user := GenerateUserPresenter(validCode.User.ID, validCode.User.Name, validCode.User.Profile, validCode.User.Email, validCode.User.GoogleToken)
			presenter = GenerateAuthenticatePresenter(token, true, user)

			return &presenter, nil
		}

	}

	return nil, errors.New("code Not found")
}
