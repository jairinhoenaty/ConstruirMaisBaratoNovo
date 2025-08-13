package store_usecase

import (
	pkgstore "construir_mais_barato/app/domain/store"
	pkguser "construir_mais_barato/app/domain/user"
	pkguseruc "construir_mais_barato/app/usecase/user"
	"fmt"
)

type DeleteStoreUC struct {
	Service     pkgstore.StoreService
	ServiceUser pkguser.UserService
	ID          *uint
}

type DeleteStoreUCParams struct {
	Service     pkgstore.StoreService
	ServiceUser pkguser.UserService
}

func NewDeleteStoreUC(params DeleteStoreUCParams) DeleteStoreUC {
	return DeleteStoreUC{
		Service:     params.Service,
		ServiceUser: params.ServiceUser,
	}
}

func (uc *DeleteStoreUC) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	oldStore, err := uc.Service.FindById(*uc.ID)

	err = uc.Service.Remove(*uc.ID)
	if err != nil {
		return err
	}

	// deve remover o profissional da tabela de usuário.
	paramsUser := pkguseruc.DeleteUserWithStoreUCParams{
		Service: uc.ServiceUser,
	}
	ucUser := pkguseruc.NewDeleteUserWithStoreUC(paramsUser)
	ucUser.Name = oldStore.Name
	ucUser.Email = oldStore.Email
	err = ucUser.Execute()
	if err != nil {
		return err
	}

	return nil
}
