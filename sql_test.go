package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(name, email, balance, rating, birth_date, married) VALUES('bud12','ai@gmail.com',10000,5.5,'1999-09-09',true)"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sukses bro !")
}

func TestSqlquery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer "
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id :", id)
		fmt.Println("name :", name)

	}
}

func TestXxxQu(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name , email, balance, rating, create_at, birth_date, married FROM customer "
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int32
		var name, email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var create_at time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &create_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("id :", id)
		if name.Valid {
			fmt.Println("name :", name.String)
		}
		if email.Valid {
			fmt.Println("email :", email.String)
		}
		fmt.Println("balance :", balance)
		fmt.Println("rating :", rating)
		if birth_date.Valid {
			fmt.Println("birth date :", birth_date.Time)
		}
		fmt.Println("create :", create_at)
		fmt.Println("married :", married)

	}

}

func TestSqlInjeksi(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "adm"

	script := "SELECT username FROM user WHERE username='" + username + "' AND password= '" + password + "' LIMIT 1 "
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		defer rows.Close()
	}
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login", username)
	} else {
		fmt.Println("Gagal login")
	}

}

func TestSqlInjeksiSave(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	script := "SELECT username FROM user WHERE username=? AND password= ? LIMIT 1 "
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		defer rows.Close()
	}
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login", username)
	} else {
		fmt.Println("Gagal login")
	}

}

func TestSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "pass"
	password := "pass"

	script := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sukses bro !")
}

func TestSqlLastId(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "ami@gmail.com"
	text := "laalallala"

	script := "INSERT INTO last_id(email, text) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, email, text)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Sukses bro !", insertId)
}

func TestSqlStatment(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO last_id(email, text) VALUES(?, ?)"
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "ami" + strconv.Itoa(i) + "@gmail"
		text := "komentar " + strconv.Itoa((i))

		result, err := stmt.ExecContext(ctx, email, text)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("commend id", id)
	}
}

func TestTransaksi(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	script := "INSERT INTO last_id(email, text) VALUES(?, ?)"
	// do transaksi
	for i := 0; i < 10; i++ {
		email := "ami" + strconv.Itoa(i) + "@gmail"
		text := "komentar " + strconv.Itoa((i))

		res, err := tx.ExecContext(ctx, script, email, text)
		if err != nil {
			panic(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			panic(err)

		}
		fmt.Println("Comment id", id)
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
