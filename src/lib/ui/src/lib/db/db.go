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

	fmt.Println(reflect.TypeOf(db))	
	
	 
	 return db
}
func QUERY(myQueryString string){
	db := db_init()
	db.Exec(myQueryString)
	fmt.Println(myQueryString)
	defer db.Close()
}