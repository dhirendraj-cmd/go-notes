package database

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
)


const (
	host   	 = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Devilisafoot123"
	dbname   = "gonotesapi"
)


func Connection() *sql.DB {
	fmt.Println("Connection to DB.... ")

	// connStr := "user=postgres password=Devilisafoot123 dbname=gonotesapi sslmode=disable"
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)

	if err !=  nil{
		panic(err)
	}

	// defer db.Close()



	// setting pool connections
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(0)


	// testing connection
	err = db.Ping()
	if err != nil{
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db

}
