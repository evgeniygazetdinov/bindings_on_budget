package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	// "reflect"
	"log"
)

// TODO add logic check db add separate functions
func DbInit() *sql.DB {
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
func QUERY_INPLACE(myQueryString string) {
	db := DbInit()
	db.Exec(myQueryString)
	defer db.Close()
}

// move this to query sets
func listByTask(my_query string) map[int]string {
	my_tasks := make(map[int]string)
	db := DbInit()
	var (
		id   int
		name string
	)
	rows, _ := db.Query(my_query)

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

func GetAllExistsTask() map[int]string {
	return listByTask("select tasks_id, name from tasks;")
}

// MOVE TO CONST AND GET THIS UP
func GetLastTask() map[int]string {
	return listByTask("select tasks_id, name from tasks order by tasks_id desc limit 1;")
}

func TaskAlreadyExistITaskMap(taskMap map[int]string, taskForCheckIn string) bool {
	for _, v := range taskMap {
		if v == taskForCheckIn {
			return true
		}
	}
	return false
}

// OUTPUT FUNCTIONS

func TASK_EXIST_IN_DB(taskForCheck string) bool {
	return TaskAlreadyExistITaskMap(GetAllExistsTask(), taskForCheck)
}
func GET_LAST_TASK() string {
	lastTask := GetLastTask()
	fmt.Println(lastTask)
	if len(lastTask) > 0 {
		for _, taskName := range lastTask {
			return taskName
		}

	}
	return " "
}
