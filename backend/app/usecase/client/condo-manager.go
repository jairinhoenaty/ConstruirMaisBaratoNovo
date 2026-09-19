package client_usecase

import (
	"strings"

	pkguser "construir_mais_barato/app/domain/user"
)

// markCondoManagers preenche IsCondoManager a partir da tabela de usuários,
// onde o flag mora. Cliente e usuário só se ligam por e-mail, então é uma
// consulta por lista em vez de um JOIN — users também tem name e created_at,
// e o JOIN deixaria ambíguas as ordenações que as listas já usam.
func markCondoManagers(userService pkguser.UserService, presenters []ClientPresenter) error {
	if userService == nil || len(presenters) == 0 {
		return nil
	}

	emails := make([]string, 0, len(presenters))
	for _, presenter := range presenters {
		emails = append(emails, presenter.Email)
	}

	found, err := userService.FindCondoManagerEmails(emails)
	if err != nil {
		return err
	}

	managers := make(map[string]bool, len(found))
	for _, email := range found {
		managers[strings.ToLower(email)] = true
	}
	for i := range presenters {
		presenters[i].IsCondoManager = managers[strings.ToLower(presenters[i].Email)]
	}
	return nil
}
