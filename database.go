package godatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/new_sql2?parseTime=true")

	if err != nil {
		panic(err)
	}
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db

}
