package main

import (
	"construir_mais_barato/cmd/app"
	"construir_mais_barato/domain/entity"
	pkgdatabase "construir_mais_barato/infra/database/mysql-db"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Carrega .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Lê variáveis do .env
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUsername == "" {
		fmt.Println("===== NÃO leu as variáveis de ambiente")
	} else {
		fmt.Println("===== conseguiu ler as variáveis de ambiente")
	}

	fmt.Println("DB_USERNAME:", dbUsername)
	fmt.Println("DB_PASSWORD:", dbPassword)
	fmt.Println("DB_HOST:", dbHost)
	fmt.Println("DB_PORT:", dbPort)
	fmt.Println("DB_NAME:", dbName)

	// Monta parâmetros
	params := &pkgdatabase.ConfigParams{
		DBUsername: dbUsername,
		DBPassword: dbPassword,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBName:     dbName,
	}

	// Faz conexão com banco
	db := pkgdatabase.ConnectionDB(params)

	
	db.AutoMigrate(&entity.Job{})

	// Inicia a aplicação
	app.Start(db)
}
