package godatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDataBase(t *testing.T) {
}

func TestDatabase(t *testing.T) {
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/new_sql2")

	if err != nil {
		panic(err)
	}
	defer db.Close()
}
