package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "symbolTables.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	myString := ""
	//for i := 0; i <= 29; i++ {
	for i := 0; i <= 29; i++ {
		str := fmt.Sprintf(`    CREATE TABLE T%d (
		Id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
        Name  TEXT,
        AltName TEXT,
        Url TEXT,
		Year INTEGER
    );

`, i)
		myString = myString + str
	}
	fmt.Println(myString)
	result, err := db.Exec(myString)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id последнего добавленного объекта
	//   fmt.Println(result.RowsAffected())  // количество добавленных строк

}
