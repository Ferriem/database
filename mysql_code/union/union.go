package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:04271017@tcp(localhost:3306)/tmp")
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

func query(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM userinfo")
	CheckErr(err)
	defer rows.Close()
	fmt.Println("userinfo table")
	fmt.Println("--------------")
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err := rows.Scan(&uid, &username, &department, &created)
		CheckErr(err)
		fmt.Println(uid, username, department, created)
	}
	fmt.Println("--------------")
}

func insert(db *sql.DB, username string, department string) int64 {
	stmt, err := db.Prepare("INSERT userinfo SET username=?, department=?, created=?")
	CheckErr(err)
	res, err := stmt.Exec(username, department, time.Now().Format("2006-01-02 15:04:05"))
	CheckErr(err)
	id, err := res.LastInsertId()
	CheckErr(err)
	query(db)
	return id
}

func update(db *sql.DB, id int64, username string) {
	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	CheckErr(err)
	_, err = stmt.Exec(username, id)
	CheckErr(err)
	query(db)
}

func delete(db *sql.DB, id int64) {
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	CheckErr(err)
	_, err = stmt.Exec(id)
	CheckErr(err)
	query(db)
}

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	fmt.Println(time.Now())
	db := connect()
	defer db.Close()
	// CREATE TABLE `userinfo` (
	//     `uid` INT(10) NOT NULL AUTO_INCREMENT,
	//     `username` VARCHAR(64) NULL DEFAULT NULL,
	//     `department` VARCHAR(64) NULL DEFAULT NULL,
	//     `created` DATE NULL DEFAULT NULL,
	//     PRIMARY KEY (`uid`)
	// );
	id := insert(db, "astaxie", "研发部")
	update(db, id, "astaxieupdate")
	delete(db, id)
}
