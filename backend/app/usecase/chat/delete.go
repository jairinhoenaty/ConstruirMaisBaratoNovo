package chat_usecase

import (
	pkgchat "construir_mais_barato/app/domain/chat"
	"fmt"
)

type DeleteChatUC struct {
	Service pkgchat.ChatService
	ID      *uint
}

type DeleteChatUCParams struct {
	Service pkgchat.ChatService
}

func NewDeleteChatUC(params DeleteChatUCParams) DeleteChatUC {
	return DeleteChatUC{
		Service: params.Service,
	}
}

func (uc *DeleteChatUC) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	err := uc.Service.Remove(*uc.ID)
	if err != nil {
		return err
	}
	return nil
}
