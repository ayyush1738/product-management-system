package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sqlx.DB

func InitDatabase() {
	dsn := "host=localhost user=postgres password=1738 dbname=product_db sslmode=disable"
	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}
	log.Println("Connected to database successfully!")
}
