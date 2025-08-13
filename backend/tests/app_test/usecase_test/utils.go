package usecase_teste

import (
	"gorm.io/gorm"

	pkgdatabase "construir_mais_barato/infra/database/mysql-db"
)

func setupTestDB() *gorm.DB {
	return pkgdatabase.ConnectionDB(nil)
}
