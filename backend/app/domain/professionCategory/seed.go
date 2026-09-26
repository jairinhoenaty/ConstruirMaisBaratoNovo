package professionCategory

import (
	"errors"

	"gorm.io/gorm"
)

const DefaultCategoryName = "Casa e Lar"

// SeedDefaultCategory garante a categoria padrão e adota nela toda profissão
// ainda sem categoria — inclusive as que já existiam antes deste recurso.
//
// Escreve em `professions` por tabela, e não pelo domínio, porque o pacote
// profession já depende deste: importá-lo de volta fecharia um ciclo.
func SeedDefaultCategory(db *gorm.DB) error {
	var category ProfessionCategory

	err := db.Where("name = ?", DefaultCategoryName).First(&category).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		category = ProfessionCategory{
			Name:        DefaultCategoryName,
			Description: "Serviços realizados no endereço do cliente",
			Active:      true,
		}
		if err := db.Create(&category).Error; err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return db.Table("professions").
		Where("category_id IS NULL AND deleted_at IS NULL").
		Update("category_id", category.ID).Error
}
