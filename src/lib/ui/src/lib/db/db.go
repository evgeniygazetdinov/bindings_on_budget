package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"fmt"
)


	// TODO add logic check db add separate functions
func DB_RUNNER(){
	db, err := sql.Open("sqlite3", "store.db")
    if err != nil {
        panic(err)
    }
    
	result, err := db.Exec("CREATE TABLE if not exists  contacts (contact_id INTEGER PRIMARY KEY,first_name TEXT NOT NULL,last_name TEXT NOT NULL,email TEXT NOT NULL UNIQUE,phone TEXT NOT NULL UNIQUE);")
	if err != nil {
        panic(err)
    }
	fmt.Println(result)

	 // количество добавленных строк
	 defer db.Close()
}