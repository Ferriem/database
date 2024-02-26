package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:04271017@tcp(localhost:3306)/classicmodels")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to the database\n")
	return db
}

func insert(db *sql.DB, Insert_information string) {
	insert, err := db.Query(Insert_information)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Inserted into subscribers table")
}

func query(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM subscribers")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	fmt.Println("subscribers table")
	fmt.Println("--------------")
	for rows.Next() {
		var id int
		var email string
		err := rows.Scan(&id, &email)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id, email)
	}
	fmt.Println("--------------")
}

func main() {
	db := connect()
	defer db.Close()
	// | subscribers | CREATE TABLE `subscribers` (
	// 	`id` int NOT NULL AUTO_INCREMENT,
	// 	`email` varchar(130) NOT NULL,
	// 	PRIMARY KEY (`id`),
	// 	UNIQUE KEY `email` (`email`)
	//   ) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |
	Insert_information := "INSERT INTO subscribers (email) VALUES ('john.doe@gmail.com') ON DUPLICATE KEY UPDATE email = 'john.doe@gmail.com';"
	insert(db, Insert_information)
	fmt.Println("Inserted into subscribers table")
	query(db)
}
