package solicitation

import (
	pkgsolicitation "construir_mais_barato/app/domain/solicitationAPP"
	"fmt"
)

type UpdateFeedbackUC struct {
	Service   pkgsolicitation.SolicitationAppService
	Assembler *UpdateFeedbackAssembler
}

type UpdateFeedbackUCParams struct {
	Service pkgsolicitation.SolicitationAppService
}

func NewUpdateFeedbackUC(params UpdateFeedbackUCParams) UpdateFeedbackUC {
	return UpdateFeedbackUC{
		Service: params.Service,
	}
}

func (uc *UpdateFeedbackUC) Execute() (*pkgsolicitation.SolicitationApp, error) {
	if uc.Assembler == nil {
		fmt.Println("UpdateFeedbackUC Assembler vazio, verifique")
		return nil, fmt.Errorf("UpdateFeedbackUC Assembler vazio, verifique")
	}

	if uc.Assembler.IdFirebase == "" {
		return nil, fmt.Errorf("idFirebase obrigatorio para gravar feedback")
	}

	result, err := uc.Service.UpdateFeedback(
		uc.Assembler.IdFirebase,
		uc.Assembler.Rating,
		uc.Assembler.Feedback,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
