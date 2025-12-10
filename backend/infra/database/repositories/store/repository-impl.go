package store_repository_impl

import (
	pkgstore "construir_mais_barato/app/domain/store"
	"encoding/json"
	"errors"
	"strings"

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
func (r *repository) FindByCategoryAndSubCategory(categoryID, cityID int, subCategories []int) ([]*pkgstore.Store, error) {
	stores := make([]*pkgstore.Store, 0)

	query := r.DB.
		Preload("City").
		Where("category_product_id = ?", categoryID).
		Where("city_id = ?", cityID)

	// Se há subcategorias solicitadas, verificar se a loja tem PELO MENOS UMA delas
	if len(subCategories) > 0 {
		// Construir condição OR para cada subcategoria
		// Exemplo: (JSON_CONTAINS(sub_categories, '1') OR JSON_CONTAINS(sub_categories, '2'))
		orConditions := make([]string, 0, len(subCategories))
		args := make([]interface{}, 0, len(subCategories))

		for _, subCat := range subCategories {
			// Cada subcategoria é um elemento individual
			subCatJSON, err := json.Marshal(subCat)
			if err != nil {
				return nil, err
			}
			orConditions = append(orConditions, "JSON_CONTAINS(sub_categories, ?)")
			args = append(args, string(subCatJSON))
		}

		// Combinar com OR: (cond1 OR cond2 OR cond3)
		whereClause := "(" + strings.Join(orConditions, " OR ") + ")"
		query = query.Where(whereClause, args...)
	}

	if err := query.Find(&stores).Error; err != nil {
		return nil, err
	}

	return stores, nil
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

// CountByCategory conta quantas lojas usam uma categoria específica
// Verifica TANTO category_product_id (categoria principal)
// QUANTO sub_categories (array JSON de subcategorias)
func (r *repository) CountByCategory(categoryID uint) (int64, error) {
	var count int64

	// Preparar JSON para busca
	categoryJSON, err := json.Marshal(categoryID)
	if err != nil {
		return 0, err
	}

	// Contar lojas que usam esta categoria como:
	// 1. Categoria principal (category_product_id) OU
	// 2. Dentro do array de subcategorias (sub_categories)
	if err := r.DB.Model(&pkgstore.Store{}).
		Where("category_product_id = ? OR JSON_CONTAINS(sub_categories, ?)", categoryID, string(categoryJSON)).
		Where("stores.deleted_at IS NULL").
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// CountBySubCategory conta quantas lojas usam uma subcategoria específica
func (r *repository) CountBySubCategory(subCategoryID uint) (int64, error) {
	var count int64

	// Construir JSON para busca (subcategoria como número individual)
	subCatJSON, err := json.Marshal(subCategoryID)
	if err != nil {
		return 0, err
	}

	if err := r.DB.Model(&pkgstore.Store{}).
		Where("JSON_CONTAINS(sub_categories, ?)", string(subCatJSON)).
		Where("stores.deleted_at IS NULL").
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// MigrateCategoryBulk migra todas as lojas de uma categoria para outra
// Processa TANTO category_product_id QUANTO sub_categories
func (r *repository) MigrateCategoryBulk(fromCategoryID, toCategoryID uint) (int64, error) {
	var totalAffected int64

	// Parte 1: Migrar lojas que usam como categoria principal (category_product_id)
	resultMain := r.DB.Model(&pkgstore.Store{}).
		Where("category_product_id = ?", fromCategoryID).
		Where("stores.deleted_at IS NULL").
		Update("category_product_id", toCategoryID)

	if resultMain.Error != nil {
		return 0, resultMain.Error
	}
	totalAffected += resultMain.RowsAffected

	// Parte 2: Migrar lojas que têm a categoria no array sub_categories
	fromCatJSON, err := json.Marshal(fromCategoryID)
	if err != nil {
		return totalAffected, err
	}

	var stores []*pkgstore.Store
	if err := r.DB.
		Where("JSON_CONTAINS(sub_categories, ?)", string(fromCatJSON)).
		Where("stores.deleted_at IS NULL").
		Find(&stores).Error; err != nil {
		return totalAffected, err
	}

	// Atualizar cada loja que tem a categoria no array
	for _, store := range stores {
		subCategories := store.SubCategories

		// Remover categoria antiga e adicionar nova
		newSubCategories := make(pkgstore.UintSlice, 0)
		for _, subCat := range subCategories {
			if subCat != fromCategoryID {
				newSubCategories = append(newSubCategories, subCat)
			}
		}

		// Adicionar nova categoria se ainda não existir
		hasNew := false
		for _, subCat := range newSubCategories {
			if subCat == toCategoryID {
				hasNew = true
				break
			}
		}
		if !hasNew {
			newSubCategories = append(newSubCategories, toCategoryID)
		}

		// Atualizar no banco
		if err := r.DB.Model(&pkgstore.Store{}).
			Where("id = ?", store.ID).
			Update("sub_categories", newSubCategories).Error; err != nil {
			continue
		}

		totalAffected++
	}

	return totalAffected, nil
}

// MigrateSubCategoryBulk migra todas as lojas de uma subcategoria para outra
// Remove a subcategoria antiga do array JSON e adiciona a nova
func (r *repository) MigrateSubCategoryBulk(fromSubCategoryID, toSubCategoryID uint) (int64, error) {
	// Buscar todas as lojas que usam a subcategoria antiga
	fromSubCatJSON, err := json.Marshal(fromSubCategoryID)
	if err != nil {
		return 0, err
	}

	var stores []*pkgstore.Store
	if err := r.DB.
		Where("JSON_CONTAINS(sub_categories, ?)", string(fromSubCatJSON)).
		Where("stores.deleted_at IS NULL").
		Find(&stores).Error; err != nil {
		return 0, err
	}

	// Atualizar cada loja
	var affected int64
	for _, store := range stores {
		// Trabalhar diretamente com o UintSlice
		subCategories := store.SubCategories

		// Remover antiga
		newSubCategories := make(pkgstore.UintSlice, 0)
		for _, subCat := range subCategories {
			if subCat != fromSubCategoryID {
				newSubCategories = append(newSubCategories, subCat)
			}
		}

		// Adicionar nova (se ainda não existir)
		hasNew := false
		for _, subCat := range newSubCategories {
			if subCat == toSubCategoryID {
				hasNew = true
				break
			}
		}
		if !hasNew {
			newSubCategories = append(newSubCategories, toSubCategoryID)
		}

		// Atualizar no banco
		if err := r.DB.Model(&pkgstore.Store{}).
			Where("id = ?", store.ID).
			Update("sub_categories", newSubCategories).Error; err != nil {
			continue
		}

		affected++
	}

	return affected, nil
}
