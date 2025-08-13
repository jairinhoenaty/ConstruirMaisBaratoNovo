package professional_usecase

type ProfessionalAssembler struct {
	ID            uint   `json:"oid,omitempty"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Telephone     string `json:"telephone"`
	LgpdAceito    string `json:"lgpdAceito"`
	Password      string `json:"password"`
	CityID        uint   `json:"cityId"`
	ProfessionIDs []uint `json:"professionIds"`
	Cep           string `json:"cep"`
	Street        string `json:"street"`
	Neighborhood  string `json:"neighborhood"`
	Image         []byte `json:"image"`
	Verified      *bool  `json:"verified"`
	OnLine        *bool  `json:"online"`
}

type CityAssembler struct {
	ID   uint   `json:"oid,omitempty"`
	Name string `json:"nome"`
	UF   string `json:"uf"`
}

type ProfissionAssembler struct {
	ID          uint   `json:"oid,omitempty"`
	Name        string `json:"nome"`
	Description string `json:"descricao"`
	Icon        string `json:"icone"`
}

type FindProfessionalByCityAndProfessionAssembler struct {
	Name         string `json:"name"`
	Limit        int    `json:"limit"`
	CityId       uint   `json:"cityId"`
	Offset       int    `json:"offset"`
	PageSize     int    `json:"pageSize"`
	ProfessionId uint   `json:"professionId"`
}

type FindByProfessionAndLocationAssembler struct {
	Limit        int     `json:"limit"`
	Offset       int     `json:"offset"`
	ProfessionId uint    `json:"professionId"`
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
	Distance     int     `json:"distance"`
}

type FindLastProfessionalsRequest struct {
	QuantityRecords int `json:"quantity"`
}

type FindWithPaginationProfessionalAssembler struct {
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
	Filter       string `json:"filter"`
	Uf           string `json:"uf"`
	ProfessionId int    `json:"profession_id"`
	Order        string `json:"order"`
}

type FindByNameAssembler struct {
	Name string `json:"name"`
}

type FindByNameAndCityAndProfessionAssembler struct {
	Name         string `json:"name"`
	CityId       uint   `json:"cityId"`
	ProfessionId uint   `json:"professionId"`
}

type FindWithPaginationProfessionalByStateAssembler struct {
	UF     string `json:"state"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}
