package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

// TODO add logic check db add separate functions
func DB_RUNNER(){
	db, _ := sql.Open("sqlite3", "./my_db.sql")
	// statement, _ := db.Prepare("create table if not exists people(id integer primary key, name text ");
	// statement.Exec()
	// defer 
	// statement, _ = db.Exec("insert into people(name) values(?)","my_name");
	db, err := sql.Open("sqlite3", "store.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

	result, err := db.Exec("insert into products (model, company, price) values ('iPhone X', $1, $2)", 
        "Apple", 72000)
    if err != nil{
        panic(err)
    }
    fmt.Println(db.LastInsertId())  // id последнего добавленного объекта
    fmt.Println(db.RowsAffected())  // количество добавленных строк
}