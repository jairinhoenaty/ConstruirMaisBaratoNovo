package solicitation

import (
	pkgsolicitation "construir_mais_barato/app/domain/solicitationAPP"
	"fmt"
)

type SaveSolicitationUC struct {
	Service   pkgsolicitation.SolicitationAppService
	Assembler *SaveSolicitationAssembler
}

type SaveSolicitationUCParams struct {
	Service pkgsolicitation.SolicitationAppService
}

func NewSaveSolicitationUC(params SaveSolicitationUCParams) SaveSolicitationUC {
	return SaveSolicitationUC{
		Service: params.Service,
	}
}

func (uc *SaveSolicitationUC) Execute() (*pkgsolicitation.SolicitationApp, error) {
	if uc.Assembler == nil {
		fmt.Println("SaveSolicitationUC Assembler vazio, verifique")
		return nil, fmt.Errorf("SaveSolicitationUC Assembler vazio, verifique")
	}
	solicitation := pkgsolicitation.SolicitationApp{
		ClientId:       uc.Assembler.ClientId,
		ClientName:     uc.Assembler.ClientName,
		Description:    uc.Assembler.Description,
		Address:        uc.Assembler.Address,
		Latitude:       uc.Assembler.Latitude,
		Longitude:      uc.Assembler.Longitude,
		ProfessionId:   uc.Assembler.ProfessionId,
		Status:         uc.Assembler.Status,
		IdFirebase:     uc.Assembler.IdFirebase,
		ProfessionalId: uc.Assembler.ProfessionalId,
		ProposalValue:  uc.Assembler.ProposalValue,
		Distance:       uc.Assembler.Distance,
	}

	result, err := uc.Service.Save(solicitation)
	if err != nil {
		return nil, err
	}
	return result, nil
}
