package chat_usecase

import (
	pkgchat "construir_mais_barato/app/domain/chat"
)

type FindByUserChatUC struct {
	Service pkgchat.ChatService
}

type FindByUserChatUCParams struct {
	Service pkgchat.ChatService
}

func NewFindByUserUC(params FindByUserChatUCParams) FindByUserChatUC {
	return FindByUserChatUC{
		Service: params.Service,
	}
}

func (uc *FindByUserChatUC) Execute(limit int, offset int, professionalID int, clienteID int) (*[]ChatPresenter, int64, error) {

	chats, total, err := uc.Service.FindByUser(limit, offset, professionalID, clienteID)

	if err != nil {
		return nil, 0, err
	}
	presenters := make([]ChatPresenter, 0)
	if len(chats) > 0 {
		for _, chat := range chats {
			presenters = append(presenters, ChatPresenter{
				ID:             chat.ID,
				Message:        chat.Message,
				ProfessionalID: chat.ProfessionalID,
				Professional:   chat.Professional,
				ClientID:       chat.ClientID,
				Client:         chat.Client,
				Origem:         chat.Origem,
				CreatedAt:      chat.CreatedAt,
			})
		}
	}
	return &presenters, total, nil
}
