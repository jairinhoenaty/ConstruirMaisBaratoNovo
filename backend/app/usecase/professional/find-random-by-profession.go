package professional_usecase

import (
    pkgprofessional "construir_mais_barato/app/domain/professional"
)

type FindRandomUC struct {
    Service   pkgprofessional.ProfessionalService
    Assembler *FindRandomAssembler
}

type FindRandomUCParams struct {
    Service pkgprofessional.ProfessionalService
}

type FindRandomAssembler struct {
    // Você pode receber por body (POST) ou via query (GET -> controller converte)
    ProfessionID   *uint  `json:"professionId"`
    ProfessionName string `json:"profession"` // opcional; usa LIKE
    Verified       *bool  `json:"verified"`
    Online         *bool  `json:"online"`
    Seed           *int64 `json:"seed"`
    Limit          int    `json:"limit"`
    Offset         int    `json:"offset"`
}

func NewFindRandomUC(params FindRandomUCParams) *FindRandomUC {
    return &FindRandomUC{
        Service: params.Service,
    }
}

func (uc *FindRandomUC) Execute() ([]*ProfessionalPresenter, int64, error) {
    var namePtr *string
    if uc.Assembler != nil && uc.Assembler.ProfessionName != "" {
        n := uc.Assembler.ProfessionName
        namePtr = &n
    }

    pros, total, err := uc.Service.FindRandom(
        uc.Assembler.ProfessionID,
        namePtr,
        uc.Assembler.Verified,
        uc.Assembler.Online,
        uc.Assembler.Seed,
        uc.Assembler.Limit,
        uc.Assembler.Offset,
    )
    if err != nil {
        return nil, 0, err
    }

    presenters := make([]*ProfessionalPresenter, 0, len(pros))
    for _, p := range pros {
        pp := GenerateProfessionalPresenter(p)
        presenters = append(presenters, &pp)
    }
    return presenters, total, nil
}
