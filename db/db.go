package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	var err error

	err = godotenv.Load()

	if err != nil {
		panic(err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	createTables()
}

func createTables() {
	createMidiaTable := `
		CREATE TABLE IF NOT EXISTS midia (
			matricula INTEGER NOT NULL,
			nome varchar(100) not null,
			streaming_favorito VARCHAR(25) NULL,
			freq_uso_redes_sociais DECIMAL(4,2) NULL,
			meio_principal_noticias VARCHAR(35) NULL,
			comunicacao_digital_principal VARCHAR(40) NULL,
			PRIMARY KEY(matricula),
			INDEX Midia_FKIndex1(matricula)
		)`

	_, err := DB.Exec(createMidiaTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not create midia table.")
	}
}
