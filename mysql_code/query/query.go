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

func print_tables(db *sql.DB) {
	fmt.Println("Database tables")
	fmt.Println("--------------")
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(table)
	}
	fmt.Println("--------------")
}

func print_columns(db *sql.DB, table string) {
	fmt.Println(table + " columns")
	rows, err := db.Query("SHOW COLUMNS FROM " + table)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("------------------------------------------------------------")
	defer rows.Close()
	for rows.Next() {
		var field, typeField, null, key, extra string
		var defaultField sql.NullString
		err := rows.Scan(&field, &typeField, &null, &key, &defaultField, &extra)
		if err != nil {
			panic(err.Error())
		}
		var defaultFieldStr string
		if defaultField.Valid {
			defaultFieldStr = defaultField.String
		} else {
			defaultFieldStr = "NULL"
		}
		fmt.Printf("%-25s %-15s %-4s %-3s %-15s %-5s\n", field, typeField, null, key, defaultFieldStr, extra)
	}
	fmt.Println("------------------------------------------------------------")
}

func print_columns_content(db *sql.DB, columns []string, table string) {
	fmt.Println(table + " content")
	query_string := "SELECT"
	for i := 0; i < len(columns); i++ {
		query_string += " " + columns[i]
		if i < len(columns)-1 {
			query_string += ","
		}
	}
	query_string += " FROM " + table
	rows, err := db.Query(query_string)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("------------------------------------------------------------")
	for i := 0; i < len(columns); i++ {
		fmt.Printf("%-15s", columns[i])
	}
	fmt.Println()
	fmt.Println("------------------------------------------------------------")
	defer rows.Close()
	for rows.Next() {
		var values []interface{}
		for i := 0; i < len(columns); i++ {
			var value string
			values = append(values, &value)
		}
		err := rows.Scan(values...)
		if err != nil {
			panic(err.Error())
		}
		for i := 0; i < len(columns); i++ {
			fmt.Printf("%-15s", *values[i].(*string))
		}
		fmt.Println()
	}
	fmt.Println("------------------------------------------------------------")
}

func main() {
	db := connect()
	defer db.Close()

	print_tables(db)

	fmt.Println()

	print_columns(db, "employees")

	fmt.Println()

	print_columns_content(db, []string{"lastName", "firstName"}, "employees")

}
