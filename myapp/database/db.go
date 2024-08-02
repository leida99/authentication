package database

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/myapp")
    if err != nil {
        log.Println("1")
        log.Fatal(err)
    }
    if err := DB.Ping(); err != nil {
        log.Println("2")
        log.Fatal(err)
    }
	createTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE
    );`
    
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}
