package chat_usecase

import pkgchat "construir_mais_barato/app/domain/chat"

type FindAllChatUC struct {
	Service   pkgchat.ChatService
	Assembler FindWithPaginationChatAssembler
}

type FindAllChatUCParams struct {
	Service pkgchat.ChatService
}

func NewFindAllChatUC(params FindAllChatUCParams) FindAllChatUC {
	return FindAllChatUC{
		Service: params.Service,
	}
}

func (uc *FindAllChatUC) Execute(limit int, offset int) (*[]ChatPresenter, int64, error) {

	chats, total, err := uc.Service.FindAll(limit, offset)
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
				CreatedAt:      chat.CreatedAt,
			})
		}
	}
	return &presenters, total, nil
}
