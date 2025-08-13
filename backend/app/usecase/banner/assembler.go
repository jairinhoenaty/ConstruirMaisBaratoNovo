package banner_usecase

type BannerAssembler struct {
	ID          uint   `json:"id,omitempty"`
	CityId      *uint  `json:"cityId"`
	Image       []byte `json:"image"`
	AccessLink  string `json:"accessLink"`
	Professions []uint `json:"professions"`
	Page        string `json:"page"`
	RegionId    *uint  `json:"regionId"`
}

type FindByCityIdAndProfessionIDAssembler struct {
	CityId       uint `json:"cityId"`
	ProfessionId uint `json:"professionId"`
}

type FindByPageAssembler struct {
	Page     string `json:"page"`
	CityId   *uint  `json:"cityId"`
	RegionId *uint  `json:"regionId"`
}
