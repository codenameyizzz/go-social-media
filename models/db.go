package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/partisimedia?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MySQL!")

	var hasil string
	row := DB.QueryRow("SELECT 'INI DB DARI GO'")
	row.Scan(&hasil)
	fmt.Println("Hasil uji coba DB:", hasil)

	row = DB.QueryRow("SELECT DATABASE()")
	row.Scan(&hasil)
	fmt.Println("Nama database aktif dari Go:", hasil)

}
