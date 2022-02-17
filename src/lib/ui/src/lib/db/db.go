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

func get_all_exists_task()map[int]string{
	my_tasks := make(map[int]string)
	db := db_init()
	var (
		id int
		name string
	)
	rows, _ := db.Query("select tasks_id, name from tasks;")
	
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		my_tasks[id] = name
	}
	defer rows.Close()
	return my_tasks
}

func task_already_exist_in_task_map(task_map map[int]string, task_for_check string) bool{
	for _, v := range task_map {
		if v == task_for_check{
			return true
		}
	}
	return false
}

// TODO rebamE ALL FUNC ON GOLANG STYLE(WHICH TYPING MORE PREFER IN GOLANG)
func TASK_EXIST_IN_DB(taskForCheck string )bool{
	return task_already_exist_in_task_map(get_all_exists_task(), taskForCheck)
}
func GET_LAST_TASK()string{

}