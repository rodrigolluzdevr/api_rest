package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

var db *sql.DB

// InitDB inicializa a conexão com o MySQL
func InitDB() {
	var err error

	// Pegando a string de conexão do ambiente
	dsn := os.Getenv("DB_DSN") // Exemplo: "user:password@tcp(localhost:3306)/database"
	if dsn == "" {
		log.Fatal("DB_DSN não configurado")
	}

	// Abrindo a conexão
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}

	// Verifica a conexão
	err = db.Ping()
	if err != nil {
		log.Fatalf("Banco de dados inacessível: %v", err)
	}

	fmt.Println("Conectado ao banco de dados com sucesso!")
}

// GetDB retorna a instância do banco de dados
func GetDB() *sql.DB {
	return db
}
