package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var lock = &sync.Mutex{}

type DataBase struct {
	DB *sql.DB
}

var dbSingleInstance *DataBase

func GetDBInstance() *DataBase {
	if dbSingleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbSingleInstance == nil {
			fmt.Println("Creating single instance now.")
			dbSingleInstance = new(DataBase)
			dbSingleInstance.initializeDb(
				os.Getenv("APP_DB_USERNAME"),
				os.Getenv("APP_DB_PASSWORD"),
				os.Getenv("APP_DB_NAME"),
				os.Getenv("APP_DB_PORT"),
				os.Getenv("APP_DB_HOST"),
			)
			dbSingleInstance.ensureTableExists()
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return dbSingleInstance
}

func (db *DataBase) initializeDb(user, password, dbname, port, host string) {
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(connectionString)
	var err error
	db.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func (database *DataBase) ensureTableExists() {
	if _, err := database.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS phone_numbers
(
    id SERIAL,
    country TEXT NOT NULL,
    number TEXT NOT NULL,
    code TEXT NOT NULL,
    state TEXT NOT NULL,
    CONSTRAINT phone_numbers_pkey PRIMARY KEY (id)
)`
