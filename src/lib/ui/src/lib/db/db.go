package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"fmt"
	"reflect"
)




	// TODO add logic check db add separate functions
func db_init() *sql.DB{
	db, err := sql.Open("sqlite3", "store.db")
    if err != nil {
        panic(err)
    }
    
	result, err := db.Exec("CREATE TABLE if not exists tasks (tasks_id INTEGER PRIMARY KEY,name TEXT NOT NULL);")
	if err != nil {
        panic(err)
		fmt.Println(result)
    }	 
	 return db
}
func QUERY_INPLACE(myQueryString string){
	db := db_init()
	db.Exec(myQueryString)
	defer db.Close()
}

func QUERY(myQueryString string){
	db := db_init()
	result, error :=db.Exec(myQueryString)
	return result
	defer db.Close()
}