package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
	// This works as a driver for database/sql package
)


var DB *sql.DB

func Init(){
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if(DB == nil){
		panic("DB is nil")
	}

	if err != nil {
		panic(err) 
	}
	DB.SetMaxOpenConns(10)	// at most 10 connections can be open
	DB.SetMaxIdleConns(5) // at most 5 connections must still open when no one used

	createTables()
}


func createTables (){
	createEventTable := `CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		location TEXT,
		date_time TEXT,
		user_id INTEGER
	)`

	_, err := DB.Exec(createEventTable)
	if err != nil {
		panic(err)
	}
}