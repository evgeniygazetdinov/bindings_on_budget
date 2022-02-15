package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"fmt"
	// "reflect"
	"log"
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


func get_all_exists_task(){
	db := db_init()
	// Prepare Query
	var (
		id int
		name string
	)
	rows, _ := db.Query("select tasks_id, name from tasks where tasks_id = ?", 1)
	
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	defer rows.Close()
}

func QUERY(){
	get_all_exists_task()
}