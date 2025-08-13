package client_repository_impl

import (
	pkgclient "construir_mais_barato/app/domain/client"
	"errors"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewClientRepositoryImpl(db *gorm.DB) pkgclient.ClientRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) ExportXLSX() ([]*pkgclient.Client, error) {
	var clients []*pkgclient.Client
	if err := r.DB.Preload("City").Order("name ASC").Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

/*
	func (r *repository) FindByNameAndCityAndProfession(name string, cityID, professionID uint, limit, offset int) ([]*pkgclient.Client, error) {
		var clients []*pkgclient.Client
		likePattern := "%" + name + "%"

		// Executa a consulta com LIKE, LIMIT e OFFSET
		if err := r.DB.Where("name LIKE ?", likePattern).
			Order("name ASC").
			Limit(limit).
			Offset(offset).
			Preload("City").
			Find(&clients).Error; err != nil {
			return nil, err
		}

		return clients, nil
	}
*/
func (r *repository) FindByName(name string) ([]*pkgclient.Client, error) {
	var clients []*pkgclient.Client
	likePattern := "%" + name + "%"

	if err := r.DB.Where("name LIKE ?", likePattern).
		Preload("City").
		Find(&clients).Error; err != nil {
		return nil, err
	}

	return clients, nil
}

/*
	func (r *repository) CountClientsByProfessionInCity(cityID uint) ([]pkgclient.ProfessionCount, error) {
		var professionCounts []pkgclient.ProfessionCount

		result := r.DB.Table("clients").
			Select("professions.name as profession_name, COUNT(*) as quantity").
			Joins("JOIN client_professions ON clients.id = client_professions.client_id").
			Joins("JOIN professions ON professions.id = client_professions.profession_id").
			Where("clients.city_id = ?", cityID).
			Group("client_professions.profession_id, professions.name").
			Scan(&professionCounts)

		if result.Error != nil {
			return nil, errors.New("Erro ao contar profissionais por profissão na cidade: " + result.Error.Error())
		}

		return professionCounts, nil
	}

	func (r *repository) CountClientsByState(uf string, limit, offset int) ([]pkgclient.CityClientCount, *int64, error) {
		var cities []pkgcity.City
		var result []pkgclient.CityClientCount

		// Recuperar cidades pela UF
		if err := r.DB.Where("uf = ?", uf).Find(&cities).Error; err != nil {
			return nil, nil, err
		}

		var total int64

		// Consulta para contar o número total de registros sem usar LIMIT e OFFSET
		if err := r.DB.Model(&pkgclient.Client{}).
			Joins("LEFT JOIN cities ON cities.id = clients.city_id").
			Where("cities.uf = ?", uf).
			Select("cities.id as city_id, cities.name as city_name, COALESCE(count(clients.id), 0) as client_count").
			Group("cities.id").
			Count(&total).Error; err != nil {
			return nil, &total, err // Retorna erro e o total
		}

		if err := r.DB.Model(&pkgclient.Client{}).
			Joins("LEFT JOIN cities ON cities.id = clients.city_id").
			Where("cities.uf = ?", uf).
			Select("cities.id as city_id, cities.name as city_name, COALESCE(count(clients.id), 0) as client_count").
			Group("cities.id").
			Limit(limit).
			Offset(offset).
			Scan(&result).Error; err != nil {
			return nil, nil, err
		}

		return result, &total, nil
	}

	func (r *repository) FindByCityAndProfession(cityID, professionID uint, limit, offset int) ([]*pkgclient.Client, int64, error) {
		var clients []*pkgclient.Client
		var total int64

		// Contar o total de profissionais que correspondem ao critério
		if err := r.DB.
			Joins("JOIN client_professions ON client_professions.client_id = clients.id").
			Where("clients.city_id = ? AND client_professions.profession_id = ?", cityID, professionID).
			Model(&pkgclient.Client{}).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		fmt.Println(cityID);
		fmt.Println(professionID);
		fmt.Println(limit);
		fmt.Println(offset);
		// Buscar os profissionais com paginação
		if err := r.DB.Preload("City").Preload("Professions").
			Joins("JOIN client_professions ON client_professions.client_id = clients.id").
			Where("clients.city_id = ? AND client_professions.profession_id = ?", cityID, professionID).
			Order("name ASC").
			Limit(limit).
			Offset(offset).
			Find(&clients).
			Error; err != nil {
			return nil, 0, err
		}
		fmt.Println(clients);

		return clients, total, nil
	}
*/
func (r *repository) FindLastClients(quantityRecords int) ([]pkgclient.Client, error) {
	var clients []pkgclient.Client

	result := r.DB.Preload("City").Order("created_at desc").Order("id desc").Limit(quantityRecords).Find(&clients)
	if result.Error != nil {
		return nil, errors.New("Erro ao selecionar o cliente: " + result.Error.Error())
	}

	return clients, nil
}

/*
	func (r *repository) CountClientsByProfession() ([]pkgclient.ProfessionCount, error) {
		var professionCounts []pkgclient.ProfessionCount

		result := r.DB.Table("client_professions").
			Select("professions.name as profession_name, COUNT(client_professions.profession_id) as quantity").
			Joins("JOIN professions ON professions.id = client_professions.profession_id").
			Group("client_professions.profession_id, professions.name").
			Scan(&professionCounts)

		if result.Error != nil {
			return nil, errors.New("Erro ao contar profissionais por profissão: " + result.Error.Error())
		}

		return professionCounts, nil
	}
*/
func (r *repository) FindAll(limit, offset int) ([]*pkgclient.Client, int64, error) {
	var clients []*pkgclient.Client

	var total int64

	// Contagem total de profissionais
	if err := r.DB.Model(&pkgclient.Client{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.DB.Preload("City").Order("name ASC").Limit(limit).Offset(offset).Find(&clients).Error; err != nil {
		return nil, 0, err
	}
	return clients, total, nil
}

func (r *repository) FindById(id uint) (*pkgclient.Client, error) {
	client := pkgclient.Client{}
	if err := r.DB.Preload("City").First(&client, id).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *repository) FindByEmail(email string) (*pkgclient.Client, error) {
	client := pkgclient.Client{}
	if err := r.DB.Preload("City").Where("email = ? ", email).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil

}

func (r *repository) Save(client pkgclient.Client) (*pkgclient.Client, error) {

	var existingClient pkgclient.Client

	// Verificar se o ID está presente para decidir entre atualizar ou criar
	if client.ID != 0 {
		// Tentar encontrar o profissional existente
		if err := r.DB.Where("id = ?", client.ID).First(&existingClient).Error; err != nil {
			return nil, err
		}
	}

	// Iniciar uma transação para garantir a atomicidade das operações
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		// Se o profissional existente for encontrado, atualizar o registro
		if existingClient.ID != 0 {
			// Atualizar o profissional
			if err := tx.Model(&existingClient).Updates(client).Error; err != nil {
				return err
			}
			// Manter o ID original no objeto client para as associações
			client.ID = existingClient.ID
		} else {
			// Criar um novo profissional
			if err := tx.Create(&client).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &client, nil

}

func (r *repository) Remove(id uint) error {

	if err := r.DB.Delete(&pkgclient.Client{}, id).Error; err != nil {
		return err
	}
	return nil
}
