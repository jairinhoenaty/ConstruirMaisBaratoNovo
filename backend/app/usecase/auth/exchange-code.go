package auth

import (
	pkgexchangecode "construir_mais_barato/app/domain/exchange-codes"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkguser "construir_mais_barato/app/domain/user"
	"errors"
	"time"
)

type ExchangeCodeUC struct {
	UserService         pkguser.UserService
	ProfessionalService pkgprofessional.ProfessionalService
	ExchangeCodeService pkgexchangecode.ExchangeCodeService
	Assembler           *ExchangeCodeAssembler
}

type ExchangeCodeUCParams struct {
	UserService         pkguser.UserService
	ProfessionalService pkgprofessional.ProfessionalService
	ExchangeCodeService pkgexchangecode.ExchangeCodeService
}

func NewExchangeCode(params ExchangeCodeUCParams) ExchangeCodeUC {
	return ExchangeCodeUC{
		UserService:         params.UserService,
		ProfessionalService: params.ProfessionalService,
		ExchangeCodeService: params.ExchangeCodeService,
	}
}

func (uc *ExchangeCodeUC) Execute() (*ExchangeCodePresenter, error) {

	if uc.Assembler == nil {
		return nil, errors.New("Invalid data")
	}

	presenter := ExchangeCodePresenter{}

	// pesquisar na tabela de usuário.
	user, _ := uc.UserService.FindById(uc.Assembler.UserId)
	// se não encontrar nenhum resultado, procurar na tabela de profissionais
	duration := 5 * time.Minute
	if user != nil {
		// Salvar no banco
		exchangeCode, err := uc.ExchangeCodeService.Generate(uc.Assembler.UserId, duration)
		if err != nil {
			return nil, err
		}
		presenter.Code = exchangeCode.Code
		return &presenter, nil

	}

	return nil, errors.New("User Not found")

}
