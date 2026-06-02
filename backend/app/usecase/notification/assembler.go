package notification_usecase

import pkguser "construir_mais_barato/app/domain/user"

type SendAppNotificationAssembler struct {
	Ids    []uint         `json:"ids"`
	IDType pkguser.IDType `json:"id_type"` // "user" ou "professional"
	Title  string         `json:"title"`
	Body   string         `json:"body"`
}
