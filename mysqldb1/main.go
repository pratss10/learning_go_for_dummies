package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Course struct {
	ID      string
	Details string
}

func GetRecords(db *sql.DB) {
	results, err := db.Query("Select * from Course")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var course Course
		err = results.Scan(&course.ID, &course.Details)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(course.ID, course.Details)
	}
}

func InsertRecord(db *sql.DB, ID string, Details string) {
	result, err := db.Exec("Insert into Course values(?, ?)", ID, Details)
	if err != nil {
		panic(err.Error())
	} else {
		if count, err := result.RowsAffected(); err == nil {
			fmt.Println(count, "row(s) affected")
		}
	}
}

//the update operation

//the delete operation

func main() {
	db, err := sql.Open("mysql", "mark1:password@tcp(127.0.0.1:3306)/CoursesDB")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database object created")
		GetRecords(db)
	}

	defer db.Close()
}
