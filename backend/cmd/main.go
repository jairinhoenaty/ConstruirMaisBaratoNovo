package main

import (
	"construir_mais_barato/cmd/app"
	pkgdatabase "construir_mais_barato/infra/database/mysql-db"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	if dbUsername == "" {
		fmt.Println("===== não leu as variaveis de ambiente")
		// dbUsername = "jairinhoagend"
		// dbPassword = "jairoconst@2024"
		// dbHost = "construirbarato.mysql.uhserver.com"
		// dbPort = "3306"
		// dbName = "construirbarato"
		
		//dbUsername = "root"
		//dbPassword = "root"
		//dbHost = "127.0.0.1"
		//dbPort = "3306"
		//dbName = "construirbarato"

		dbUsername = "sql10772307"
		dbPassword = "QkUizWhTa5"
		dbHost = "sql10.freesqldatabase.com"
		dbPort = "3306"
		dbName = "sql10772307"

	} else {
		fmt.Println("===== conseguiu ler as variaveis de ambiente")
	}

	/* LOCAL*/
	// dbPassword := "root"
	// dbHost := "127.0.0.1"
	// dbPort := "3306"
	// dbName := "construirbarato"

	// Imprime os valores das variáveis
	fmt.Println("DB_USERNAME:", dbUsername)
	fmt.Println("DB_PASSWORD:", dbPassword)
	fmt.Println("DB_HOST:", dbHost)
	fmt.Println("DB_PORT:", dbPort)
	fmt.Println("DB_NAME:", dbName)

	// Configuração de conexão ao banco de dados...
	params := &pkgdatabase.ConfigParams{
		DBUsername: dbUsername,
		DBPassword: dbPassword,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBName:     dbName,
	}

	db := pkgdatabase.ConnectionDB(params)

	app.Start(db)
}
