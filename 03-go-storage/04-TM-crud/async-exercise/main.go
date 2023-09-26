package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbPath := "root:@tcp(localhost:3306)/storage"

	storageDB, err := sql.Open("mysql", dbPath)

	if err != nil {
		panic(err)
	}

	defer storageDB.Close()

	if err = storageDB.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database connection succesful")

	dbStats := storageDB.Stats()
	log.Println(dbStats)
}
