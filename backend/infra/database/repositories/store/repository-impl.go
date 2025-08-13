package store_repository_impl

import (
	pkgstore "construir_mais_barato/app/domain/store"
	"errors"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewStoreRepositoryImpl(db *gorm.DB) pkgstore.StoreRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) ExportXLSX() ([]*pkgstore.Store, error) {
	var stores []*pkgstore.Store
	if err := r.DB.Preload("City").Where("stores.deleted_at IS NULL").Order("name ASC").Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}

/*
	func (r *repository) FindByNameAndCityAndProfession(name string, cityID, professionID uint, limit, offset int) ([]*pkgstore.Store, error) {
		var stores []*pkgstore.Store
		likePattern := "%" + name + "%"

		// Executa a consulta com LIKE, LIMIT e OFFSET
		if err := r.DB.Where("name LIKE ?", likePattern).
			Order("name ASC").
			Limit(limit).
			Offset(offset).
			Preload("City").
			Preload("Professions").
			Find(&stores).Error; err != nil {
			return nil, err
		}

		return stores, nil
	}
*/
func (r *repository) FindByName(name string) ([]*pkgstore.Store, error) {
	var stores []*pkgstore.Store
	likePattern := "%" + name + "%"

	if err := r.DB.Where("name LIKE ?", likePattern).
		Where("stores.deleted_at IS NULL").
		Preload("City").
		Find(&stores).Error; err != nil {
		return nil, err
	}

	return stores, nil
}

/*
	func (r *repository) CountStoresByProfessionInCity(cityID uint) ([]pkgstore.ProfessionCount, error) {
		var professionCounts []pkgstore.ProfessionCount

		result := r.DB.Table("stores").
			Select("professions.name as profession_name, COUNT(*) as quantity").
			Joins("JOIN store_professions ON stores.id = store_professions.store_id").
			Joins("JOIN professions ON professions.id = store_professions.profession_id").
			Where("stores.city_id = ?", cityID).
			Group("store_professions.profession_id, professions.name").
			Scan(&professionCounts)

		if result.Error != nil {
			return nil, errors.New("Erro ao contar profissionais por profissão na cidade: " + result.Error.Error())
		}

		return professionCounts, nil
	}

	func (r *repository) CountStoresByState(uf string, limit, offset int) ([]pkgstore.CityStoreCount, *int64, error) {
		var cities []pkgcity.City
		var result []pkgstore.CityStoreCount

		// Recuperar cidades pela UF
		if err := r.DB.Where("uf = ?", uf).Find(&cities).Error; err != nil {
			return nil, nil, err
		}

		var total int64

		// Consulta para contar o número total de registros sem usar LIMIT e OFFSET
		if err := r.DB.Model(&pkgstore.Store{}).
			Joins("LEFT JOIN cities ON cities.id = stores.city_id").
			Where("cities.uf = ?", uf).
			Select("cities.id as city_id, cities.name as city_name, COALESCE(count(stores.id), 0) as store_count").
			Group("cities.id").
			Count(&total).Error; err != nil {
			return nil, &total, err // Retorna erro e o total
		}

		if err := r.DB.Model(&pkgstore.Store{}).
			Joins("LEFT JOIN cities ON cities.id = stores.city_id").
			Where("cities.uf = ?", uf).
			Select("cities.id as city_id, cities.name as city_name, COALESCE(count(stores.id), 0) as store_count").
			Group("cities.id").
			Limit(limit).
			Offset(offset).
			Scan(&result).Error; err != nil {
			return nil, nil, err
		}

		return result, &total, nil
	}

	func (r *repository) FindByCityAndProfession(cityID, professionID uint, limit, offset int) ([]*pkgstore.Store, int64, error) {
		var stores []*pkgstore.Store
		var total int64

		// Contar o total de profissionais que correspondem ao critério
		if err := r.DB.
			Joins("JOIN store_professions ON store_professions.store_id = stores.id").
			Where("stores.city_id = ? AND store_professions.profession_id = ?", cityID, professionID).
			Model(&pkgstore.Store{}).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		fmt.Println(cityID);
		fmt.Println(professionID);
		fmt.Println(limit);
		fmt.Println(offset);
		// Buscar os profissionais com paginação
		if err := r.DB.Preload("City").Preload("Professions").
			Joins("JOIN store_professions ON store_professions.store_id = stores.id").
			Where("stores.city_id = ? AND store_professions.profession_id = ?", cityID, professionID).
			Order("name ASC").
			Limit(limit).
			Offset(offset).
			Find(&stores).
			Error; err != nil {
			return nil, 0, err
		}
		fmt.Println(stores);

		return stores, total, nil
	}
*/
func (r *repository) FindLastStores(quantityRecords int) ([]pkgstore.Store, error) {
	var professionais []pkgstore.Store

	result := r.DB.Preload("City").
		Where("stores.deleted_at IS NULL").Order("created_at desc").Order("id desc").Limit(quantityRecords).Find(&professionais)
	if result.Error != nil {
		return nil, errors.New("Erro ao selecionar o lojas: " + result.Error.Error())
	}

	return professionais, nil
}

/*
	func (r *repository) CountStoresByProfession() ([]pkgstore.ProfessionCount, error) {
		var professionCounts []pkgstore.ProfessionCount

		result := r.DB.Table("store_professions").
			Select("professions.name as profession_name, COUNT(store_professions.profession_id) as quantity").
			Joins("JOIN professions ON professions.id = store_professions.profession_id").
			Group("store_professions.profession_id, professions.name").
			Scan(&professionCounts)

		if result.Error != nil {
			return nil, errors.New("Erro ao contar profissionais por profissão: " + result.Error.Error())
		}

		return professionCounts, nil
	}
*/
func (r *repository) FindAll(limit, offset int) ([]*pkgstore.Store, int64, error) {
	var stores []*pkgstore.Store

	var total int64

	// Contagem total de profissionais
	if err := r.DB.Model(&pkgstore.Store{}).Where("stores.deleted_at IS NULL").Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.DB.Preload("City").Where("stores.deleted_at IS NULL").Order("name ASC").Limit(limit).Offset(offset).Find(&stores).Error; err != nil {
		return nil, 0, err
	}
	return stores, total, nil
}

func (r *repository) FindById(id uint) (*pkgstore.Store, error) {
	store := pkgstore.Store{}
	if err := r.DB.Preload("City").Where("stores.deleted_at IS NULL").First(&store, id).Error; err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *repository) FindByEmail(email string) (*pkgstore.Store, error) {
	store := pkgstore.Store{}
	if err := r.DB.Preload("City").Where("email = ? ", email).First(&store).Error; err != nil {
		return nil, err
	}
	return &store, nil

}

func (r *repository) Save(store pkgstore.Store) (*pkgstore.Store, error) {

	var existingStore pkgstore.Store

	// Verificar se o ID está presente para decidir entre atualizar ou criar
	if store.ID != 0 {
		// Tentar encontrar o profissional existente
		if err := r.DB.Where("id = ?", store.ID).First(&existingStore).Error; err != nil {
			return nil, err
		}
	}

	// Iniciar uma transação para garantir a atomicidade das operações
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		// Se o profissional existente for encontrado, atualizar o registro
		if existingStore.ID != 0 {
			// Atualizar o profissional
			if err := tx.Model(&existingStore).Updates(store).Error; err != nil {
				return err
			}
			// Manter o ID original no objeto store para as associações
			store.ID = existingStore.ID
		} else {
			// Criar um novo profissional
			if err := tx.Create(&store).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &store, nil

}

func (r *repository) Remove(id uint) error {

	if err := r.DB.Delete(&pkgstore.Store{}, id).Error; err != nil {
		return err
	}
	return nil
}
