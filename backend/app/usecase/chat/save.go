package chat_usecase

import pkgchat "construir_mais_barato/app/domain/chat"

type SaveChatUC struct {
	Service   pkgchat.ChatService
	Assembler *ChatAssembler
}

type SaveChatUCParams struct {
	Service pkgchat.ChatService
}

func NewSaveChatUC(params SaveChatUCParams) SaveChatUC {
	return SaveChatUC{
		Service: params.Service,
	}
}

func (uc *SaveChatUC) Execute() (*ChatPresenter, error) {
	chat := GenerateChat(uc.Assembler)
	chatSaved, err := uc.Service.Save(chat)
	if err != nil {
		return nil, err
	}
	chatPresenter := GenerateChatPresenter(chatSaved)

	return &chatPresenter, nil

}
