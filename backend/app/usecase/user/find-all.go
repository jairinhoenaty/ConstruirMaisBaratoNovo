package user_usecase

import pkguser "construir_mais_barato/app/domain/user"

type FindAllUserUC struct {
	Service pkguser.UserService
}

type FindAllUserUCParams struct {
	Service pkguser.UserService
}

func NewFindAllUserUC(params FindAllUserUCParams) FindAllUserUC {
	return FindAllUserUC{
		Service: params.Service,
	}
}

func (uc *FindAllUserUC) Execute() (*[]UserPresenter, error) {

	users, err := uc.Service.FindAll()
	if err != nil {
		return nil, err
	}
	presenters := make([]UserPresenter, 0)
	if len(users) > 0 {
		for _, user := range users {
			presenters = append(presenters, UserPresenter{
				ID:      user.ID,
				Name:    user.Name,
				Email:   user.Email,
				Profile: user.Profile,
			})
		}
	}
	return &presenters, nil
}
