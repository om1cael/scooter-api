package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("could not load .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPass, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("could not initialize db: ", err)
	}

	r := gin.Default()

	err = db.Ping()
	if err != nil {
		log.Fatal("could not ping the database: ", err)
	}

	r.Run()
}
