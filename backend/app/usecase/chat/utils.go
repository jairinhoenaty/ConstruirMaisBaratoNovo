package chat_usecase

import (
	pkgchat "construir_mais_barato/app/domain/chat"
)

func GenerateChat(assembler *ChatAssembler) pkgchat.Chat {
	chat := pkgchat.Chat{}
	if assembler != nil {
		chat.ID = assembler.ID
		chat.Message = assembler.Message
		chat.ProfessionalID = assembler.ProfessionalID
		chat.ClientID = assembler.ClientID
		chat.Origem = assembler.Origem

	}
	return chat
}

func GenerateChatPresenter(chat *pkgchat.Chat) ChatPresenter {
	presenter := ChatPresenter{}
	if chat != nil {
		presenter.ID = chat.ID
		presenter.Message = chat.Message
		presenter.ProfessionalID = chat.ProfessionalID
		presenter.Professional = chat.Professional
		presenter.ClientID = chat.ClientID
		presenter.Client = chat.Client
		presenter.Origem = chat.Origem
		presenter.CreatedAt = chat.CreatedAt

	}

	return presenter
}
