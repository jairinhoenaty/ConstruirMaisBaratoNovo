package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	pkgbanner "construir_mais_barato/app/domain/banner"
	pkgbudget "construir_mais_barato/app/domain/budget"
	pkgchat "construir_mais_barato/app/domain/chat"
	pkgcity "construir_mais_barato/app/domain/city"
	pkgclient "construir_mais_barato/app/domain/client"
	pkgcontact "construir_mais_barato/app/domain/contact"
	pkgsexchangecode "construir_mais_barato/app/domain/exchange-codes"
	pkgjob "construir_mais_barato/app/domain/job"
	pkgplan "construir_mais_barato/app/domain/plan"
	pkgproduct "construir_mais_barato/app/domain/product"
	pkgproductCategory "construir_mais_barato/app/domain/productCategory"
	pkgprofession "construir_mais_barato/app/domain/profession"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgregion "construir_mais_barato/app/domain/region"
	pkgsolicitationapp "construir_mais_barato/app/domain/solicitationAPP"
	pkgstore "construir_mais_barato/app/domain/store"
	pkgsubscription "construir_mais_barato/app/domain/subscription"
	pkgunlockprice "construir_mais_barato/app/domain/unlock-price"
	pkgunlockedbudget "construir_mais_barato/app/domain/unlocked-budget"
	pkguser "construir_mais_barato/app/domain/user"
)

type ConfigParams struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func ConnectionDB(params *ConfigParams) *gorm.DB {
	var db *gorm.DB
	var err error

	if params == nil {
		// Configurações padrão para SQLite em memória para testes
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	} else {

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			params.DBUsername, params.DBPassword, params.DBHost, params.DBPort, params.DBName)

		// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		// 	params.DBUsername, params.DBPassword, params.DBHost, params.DBName)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	}

	if err != nil {
		panic(err)
	}

	// Habilitando o modo de log (opcional)
	// db.LogMode(true)

	// Executando as migrações
	// Aqui você deve executar suas migrações
	// migrations.RunMigrations(db)

	// Executando os seeds
	// Aqui você deve executar seus seeds
	// seeds.CreateSeedData(db)

	// Criando tabelas automaticamente (opcional)
	db.AutoMigrate(&pkgcity.City{})
	db.AutoMigrate(&pkguser.User{})
	db.AutoMigrate(&pkgcontact.Contact{})
	db.AutoMigrate(&pkgchat.Chat{})
	db.AutoMigrate(&pkgprofession.Profession{})
	db.AutoMigrate(&pkgprofessional.Professional{})
	db.AutoMigrate(&pkgbudget.Budget{})
	db.AutoMigrate(&pkgbudget.BudgetRefusal{})
	db.AutoMigrate(&pkgbanner.Banner{})
	db.AutoMigrate(&pkgproduct.Product{})
	db.AutoMigrate(&pkgproductCategory.ProductCategory{})
	db.AutoMigrate(&pkgclient.Client{})
	db.AutoMigrate(&pkgstore.Store{})
	db.AutoMigrate(&pkgregion.Region{})
	db.AutoMigrate(&pkgjob.Job{})
	db.AutoMigrate(&pkgunlockedbudget.UnlockedBudget{})
	db.AutoMigrate(&pkgplan.Plan{})
	db.AutoMigrate(&pkgsubscription.Subscription{})
	db.AutoMigrate(&pkgunlockprice.UnlockPrice{})
	db.AutoMigrate(&pkgsolicitationapp.SolicitationApp{})
	db.AutoMigrate(&pkgsexchangecode.ExchangeCode{})

	// Seed de planos (popula planos iniciais se não existirem)
	if err := pkgplan.SeedPlans(db); err != nil {
		fmt.Printf("Warning: Failed to seed plans: %v\n", err)
	}

	return db
}
