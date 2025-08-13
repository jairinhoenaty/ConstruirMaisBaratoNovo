package professional_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
)

type ExportXLSXProfessionalUC struct {
	Service pkgprofessional.ProfessionalService
}

type ExportXLSXProfessionalUCParams struct {
	Service pkgprofessional.ProfessionalService
}

func NewExportXLSXProfessionalUC(params ExportXLSXProfessionalUCParams) ExportXLSXProfessionalUC {
	return ExportXLSXProfessionalUC{
		Service: params.Service,
	}
}

func (uc *ExportXLSXProfessionalUC) Execute() (*[]ProfessionalPresenter, error) {

	professionals, err := uc.Service.ExportXLSX()
	if err != nil {
		return nil, err
	}

	presenters := make([]ProfessionalPresenter, 0)
	if len(professionals) > 0 {
		for _, professional := range professionals {
			professionalPresenter := GenerateProfessionalPresenter(professional)
			presenters = append(presenters, professionalPresenter)
		}
	}

	return &presenters, nil
}
