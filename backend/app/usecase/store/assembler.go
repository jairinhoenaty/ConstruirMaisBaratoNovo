package store_usecase

type StoreAssembler struct {
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

type FindStoreByCityAndProfessionAssembler struct {
	Name         string `json:"name"`
	Limit        int    `json:"limit"`
	CityId       uint   `json:"cityId"`
	Offset       int    `json:"offset"`
	PageSize     int    `json:"pageSize"`
	ProfessionId uint   `json:"professionId"`
}

type FindLastStoresRequest struct {
	QuantityRecords int `json:"quantity"`
}

type FindWithPaginationStoreAssembler struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type FindByNameAssembler struct {
	Name string `json:"name"`
}

type FindByNameAndCityAndProfessionAssembler struct {
	Name         string `json:"name"`
	CityId       uint   `json:"cityId"`
	ProfessionId uint   `json:"professionId"`
}

type FindWithPaginationStoreByStateAssembler struct {
	UF     string `json:"state"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}
