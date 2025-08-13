package profession_repository_impl

import (
	pkgcity "construir_mais_barato/app/domain/city"
	pkgprofession "construir_mais_barato/app/domain/profession"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewProfessionalRepositoryImpl(db *gorm.DB) pkgprofessional.ProfessionalRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) ExportXLSX() ([]*pkgprofessional.Professional, error) {
	var professionals []*pkgprofessional.Professional
	if err := r.DB.Preload("City").Preload("Professions").Order("name ASC").Find(&professionals).Error; err != nil {
		return nil, err
	}
	return professionals, nil
}

func (r *repository) FindByNameAndCityAndProfession(name string, cityID, professionID uint, limit, offset int) ([]*pkgprofessional.Professional, error) {
	var professionals []*pkgprofessional.Professional
	likePattern := "%" + name + "%"

	// Executa a consulta com LIKE, LIMIT e OFFSET
	if err := r.DB.Where("name LIKE ?", likePattern).
		Where("professionals.deleted_at IS NULL").
		Order("name ASC").
		Limit(limit).
		Offset(offset).
		Preload("City").
		Preload("Professions").
		Find(&professionals).Error; err != nil {
		return nil, err
	}

	return professionals, nil
}

func (r *repository) FindByName(name string) ([]*pkgprofessional.Professional, error) {
	var professionals []*pkgprofessional.Professional
	likePattern := "%" + name + "%"

	if err := r.DB.Where("name LIKE ?", likePattern).
		Preload("City").
		Preload("Professions").
		Find(&professionals).Error; err != nil {
		return nil, err
	}

	return professionals, nil
}

func (r *repository) CountProfessionalsByProfessionInCity(cityID uint) ([]pkgprofessional.ProfessionCount, error) {
	var professionCounts []pkgprofessional.ProfessionCount

	result := r.DB.Table("professionals").
		Select("professions.name as profession_name, COUNT(*) as quantity").
		Joins("JOIN professional_professions ON professionals.id = professional_professions.professional_id").
		Joins("JOIN professions ON professions.id = professional_professions.profession_id").
		Where("professionals.city_id = ?", cityID).
		Where("professionals.deleted_at IS NULL").
		Group("professional_professions.profession_id, professions.name").
		Scan(&professionCounts)

	if result.Error != nil {
		return nil, errors.New("Erro ao contar profissionais por profissão na cidade: " + result.Error.Error())
	}

	return professionCounts, nil
}

func (r *repository) CountProfessionalsByState(uf string, limit, offset int) ([]pkgprofessional.UFProfessionalCount, *int64, error) {
	//var cities []pkgcity.City
	var result []pkgprofessional.UFProfessionalCount

	// Recuperar cidades pela UF
	//if err := r.DB.Where("uf = ?", uf).Find(&cities).Error; err != nil {
	//	return nil, nil, err
	//}

	var total int64

	// Consulta para contar o número total de registros sem usar LIMIT e OFFSET
	if err := r.DB.Model(&pkgprofessional.Professional{}).
		Joins("LEFT JOIN cities ON cities.id = professionals.city_id").
		//Where("cities.uf = ?", uf).
		Select("cities.uf as uf_name, COALESCE(count(professionals.id), 0) as professional_count").
		Where("professionals.deleted_at IS NULL").
		Group("cities.uf").
		Count(&total).Error; err != nil {
		//return nil, &total, err // Retorna erro e o total
	}

	if err := r.DB.Model(&pkgprofessional.Professional{}).
		Joins("LEFT JOIN cities ON cities.id = professionals.city_id").
		//Where("cities.uf = ?", uf).
		Select("cities.uf as uf_name, COALESCE(count(professionals.id), 0) as professional_count").
		Where("professionals.deleted_at IS NULL").
		Group("cities.uf").
		Limit(limit).
		Offset(offset).
		Scan(&result).Error; err != nil {
		//return result, &total, err
	}

	return result, &total, nil
}

func (r *repository) CountCityProfessionalsByState(uf string, limit, offset int) ([]pkgprofessional.CityProfessionalCount, *int64, error) {
	var cities []pkgcity.City
	var result []pkgprofessional.CityProfessionalCount

	// Recuperar cidades pela UF
	if err := r.DB.Where("uf = ?", uf).Find(&cities).Error; err != nil {
		return nil, nil, err
	}

	var total int64

	// Consulta para contar o número total de registros sem usar LIMIT e OFFSET
	if err := r.DB.Model(&pkgprofessional.Professional{}).
		Joins("LEFT JOIN cities ON cities.id = professionals.city_id").
		Where("cities.uf = ?", uf).
		Select("cities.id as city_id, cities.name as city_name, COALESCE(count(professionals.id), 0) as professional_count").
		Where("professionals.deleted_at IS NULL").
		Count(&total).Error; err != nil {
		return nil, &total, err // Retorna erro e o total
	}

	if err := r.DB.Model(&pkgprofessional.Professional{}).
		Joins("LEFT JOIN cities ON cities.id = professionals.city_id").
		Where("cities.uf = ?", uf).
		Select("cities.id as city_id, cities.name as city_name, COALESCE(count(professionals.id), 0) as professional_count").
		Where("professionals.deleted_at IS NULL").
		Group("cities.id").
		Limit(limit).
		Offset(offset).
		Scan(&result).Error; err != nil {
		return nil, nil, err
	}

	return result, &total, nil
}

func (r *repository) FindByCityAndProfession(cityID, professionID uint, limit, offset int) ([]*pkgprofessional.Professional, int64, error) {
	var professionals []*pkgprofessional.Professional

	// Buscar os profissionais com paginação
	if err := r.DB.Preload("City").Preload("Professions").
		Joins("JOIN professional_professions ON professional_professions.professional_id = professionals.id").
		Where("professionals.city_id = ? AND professional_professions.profession_id = ?", cityID, professionID).
		Where("professionals.deleted_at IS NULL").
		Order("verified DESC,name ASC").
		Limit(limit).
		Offset(offset).
		Find(&professionals).
		Error; err != nil {
		return nil, 0, err
	}

	return professionals, int64(len(professionals)), nil
}

/*
	func (r *repository) FindByCityAndProfession(cityID, professionID uint, verified *bool, limit, offset int) ([]*pkgprofessional.Professional, int64, error) {
		var professionals []*pkgprofessional.Professional
		//var total int64
		var where string = ""

		if verified != nil {

			where = fmt.Sprintf("professionals.verified = %t", *verified) + " and professionals.on_line=true"
		}


		// Buscar os profissionais com paginação
		if err := r.DB.
			Select("professionals.*,(6371*acos(cos(radians(-22.3968592))*cos(radians(latitude))"+
				"*cos( radians( longitude ) - radians(-43.1236966) ) +"+
				"sin( radians(-22.3968592) ) * sin( radians( latitude ) ) ) ) as distancia").
			Preload("City").
			Preload("Professions").
			Joins("JOIN professional_professions ON professional_professions.professional_id = professionals.id").
			Where("professionals.city_id = ? AND professional_professions.profession_id = ?", cityID, professionID).
			Where("professionals.deleted_at IS NULL").
			Where(where).
			//Having("distancia <= 10").
			Order("name ASC").
			Limit(limit).
			Offset(offset).
			Find(&professionals).
			Error; err != nil {
			return nil, 0, err
		}

		return professionals, int64(len(professionals)), nil
	}
*/
func (r *repository) FindByProfessionAndLocation(professionID uint, latitude float32, longitude float32, distance, limit, offset int) ([]*pkgprofessional.Professional, int64, error) {
	var professionals []*pkgprofessional.Professional
	var where string = ""
	var latitudeStr = fmt.Sprintf("%f", latitude)
	var longitudeStr = fmt.Sprintf("%f", longitude)
	where = "professionals.verified = true and professionals.on_line=true"

	// Buscar os profissionais com paginação
	if err := r.DB.
		Select("professionals.*,(6371*acos(cos(radians("+latitudeStr+"))*cos(radians(latitude))"+
			"*cos( radians( longitude ) - radians("+longitudeStr+") ) +"+
			"sin( radians("+latitudeStr+") ) * sin( radians( latitude ) ) ) ) as distance").
		Preload("City").
		Preload("Professions").
		Joins("JOIN professional_professions ON professional_professions.professional_id = professionals.id").
		Where("professional_professions.profession_id = ?", professionID).
		Where("professionals.deleted_at IS NULL").
		Where(where).
		Having("distance <= ?", distance).
		Order("name ASC").
		Limit(limit).
		Offset(offset).
		Find(&professionals).
		Error; err != nil {
		return nil, 0, err
	}

	return professionals, int64(len(professionals)), nil
}

func (r *repository) FindLastProfessionals(quantityRecords int) ([]pkgprofessional.Professional, error) {
	var professionais []pkgprofessional.Professional

	result := r.DB.Preload("City").Preload("Professions").Where("professionals.deleted_at IS NULL").Order("created_at desc").Order("id desc").Limit(quantityRecords).Find(&professionais)
	if result.Error != nil {
		return nil, errors.New("Erro ao selecionar o profissional: " + result.Error.Error())
	}

	return professionais, nil
}

func (r *repository) CountProfessionalsByProfession() ([]pkgprofessional.ProfessionCount, error) {
	var professionCounts []pkgprofessional.ProfessionCount
	fmt.Println("Count by professions")
	result := r.DB.Table("professional_professions").
		Select("professions.name as profession_name, COUNT(professional_professions.profession_id) as quantity").
		Joins("JOIN professions ON professions.id = professional_professions.profession_id").
		Group("professional_professions.profession_id, professions.name").
		Scan(&professionCounts)

	if result.Error != nil {
		return nil, errors.New("Erro ao contar profissionais por profissão: " + result.Error.Error())
	}

	return professionCounts, nil
}

func (r *repository) FindAll(limit, offset int, filter string, uf string, professionId int, order string) ([]*pkgprofessional.Professional, int64, error) {
	var professionals []*pkgprofessional.Professional

	var total int64
	var where string = "0=0"
	var joins2 string = ""
	if filter != "" {
		where = where +
			" and (UPPER(professionals.name) like '%" + strings.ToUpper(filter) +
			"%' or UPPER(professionals.email) like '%" + strings.ToUpper(filter) + "%')"
	}
	if uf != "" {
		where = where + " and c.uf='" + uf + "'"
	}
	if professionId != 0 {
		where = where + " and pp.profession_id=" + strconv.Itoa(professionId)
		joins2 = "LEFT OUTER JOIN professional_professions pp ON professionals.id = pp.professional_id"
	}
	// Contagem total de profissionais
	if err := r.DB.Model(&pkgprofessional.Professional{}).
		Joins("JOIN cities c ON professionals.city_id = c.id").
		Joins(joins2).
		//Joins("LEFT OUTER JOIN professional_professions pp ON professionals.id = pp.professional_id").
		Where("professionals.deleted_at IS NULL").
		Where(where).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if order == "" {
		order = "professionals.name ASC"
	}

	if err := r.DB.
		Joins("JOIN cities c ON professionals.city_id = c.id").
		//Joins("LEFT OUTER JOIN professional_professions pp ON professionals.id = pp.professional_id").
		Joins(joins2).
		Preload("City").
		Preload("Professions").
		Where("professionals.deleted_at IS NULL").
		Where(where).
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&professionals).Error; err != nil {
		return nil, 0, err
	}

	return professionals, total, nil
}

func (r *repository) FindById(id uint) (*pkgprofessional.Professional, error) {
	professional := pkgprofessional.Professional{}
	if err := r.DB.Preload("City").Preload("Professions").First(&professional, id).Error; err != nil {
		return nil, err
	}
	return &professional, nil
}

func (r *repository) FindByEmail(email string) (*pkgprofessional.Professional, error) {
	professional := pkgprofessional.Professional{}
	if err := r.DB.Preload("City").Preload("Professions").Where("email = ? ", email).First(&professional).Error; err != nil {
		return nil, err
	}
	print("PASSOU finfbyemail")
	return &professional, nil

}

func (r *repository) Save(professional pkgprofessional.Professional) (*pkgprofessional.Professional, error) {

	var existingProfessional pkgprofessional.Professional

	// Verificar se o ID está presente para decidir entre atualizar ou criar
	if professional.ID != 0 {
		// Tentar encontrar o profissional existente
		if err := r.DB.Where("id = ?", professional.ID).First(&existingProfessional).Error; err != nil {
			return nil, err
		}
	}
	fmt.Println("ID:", professional.ID)
	fmt.Println("existing ID:", existingProfessional.ID)
	fmt.Println("verified update:", professional.Verified)
	fmt.Println("verified exist:", existingProfessional.Verified)

	// Iniciar uma transação para garantir a atomicidade das operações
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		// Se o profissional existente for encontrado, atualizar o registro
		if existingProfessional.ID != 0 {
			// Atualizar o profissional
			if err := tx.Model(&existingProfessional).Updates(professional).Error; err != nil {
				return err
			}
			// Manter o ID original no objeto professional para as associações
			professional.ID = existingProfessional.ID
		} else {
			// Criar um novo profissional
			if err := tx.Create(&professional).Error; err != nil {
				return err
			}
		}

		// Limpar associações existentes (opcional, dependendo dos requisitos)
		if err := tx.Model(&professional).Association("Professions").Clear(); err != nil {
			return err
		}

		// Adicionar novas associações
		var professions []pkgprofession.Profession
		if len(professional.ProfessionIDs) > 0 {
			if err := tx.Where("id IN ?", professional.ProfessionIDs).Find(&professions).Error; err != nil {
				return err
			}
			if err := tx.Model(&professional).Association("Professions").Replace(&professions); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &professional, nil

}

func (r *repository) Remove(id uint) error {

	if err := r.DB.Table("professional_professions").
		Where("professional_id = ?", id).
		Delete(nil).Error; err != nil {
		return err
	}
	if err := r.DB.Delete(&pkgprofessional.Professional{}, id).Error; err != nil {
		return err
	}
	return nil
}
