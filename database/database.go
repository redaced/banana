package database

import (
	"database/sql"
	"fmt"
	// "log"
	_ "github.com/go-sql-driver/mysql"
)

//connecting to mysql database
func Opendb(username, password, databaseName string) *sql.DB {
	config := fmt.Sprintf("%s:%s@/%s", username, password, databaseName)
	db, error := sql.Open("mysql", config)
	if error != nil {
		fmt.Println("Can not connect to database, check your credentials")
		return nil
	}
	return db
}

// //checking if there was an error during function execution
// func Error_check(err error) {
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

// //Inserting user informations to database table
// func InsertUserEntity(username, password string, db *sql.DB) {
// 	command, error := db.Prepare("INSERT INTO tbl_user(name, password) VALUES(?, ?)")
// 	if error != nil {
// 		log.Fatal(error)
// 	}
// 	res, error := command.Exec(username, password)
// 	lastId, error := res.LastInsertId()
// 	fmt.Println(lastId)
// }