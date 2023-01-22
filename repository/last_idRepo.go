package repository

import (
	"context"
	"database/sql"
	"errors"

	// "errors"
	"go-database/entity"
	"strconv"
	// "golang.org/x/crypto/openpgp/errors"
)

type last_IdRepo struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Last_IdRepository {
	return &last_IdRepo{DB: db}
}

func (repo *last_IdRepo) Insert(ctx context.Context, last_Id entity.Last_Id) (entity.Last_Id, error) {
	script := "INSERT INTO last_id(email, text) VALUES (?, ?)"
	res, err := repo.DB.ExecContext(ctx, script, last_Id.Email, last_Id.Text)
	if err != nil {
		return last_Id, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return last_Id, err
	}
	last_Id.Id = int32(id)
	return last_Id, nil
}
func (repo *last_IdRepo) FindById(ctx context.Context, id int32) (entity.Last_Id, error) {
	script := "SELECT id, email, text FROM last_id WHERE id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, script, id)
	last_id := entity.Last_Id{}
	if err != nil {
		return last_id, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&last_id.Id, &last_id.Email, &last_id.Text)
		return last_id, nil
	} else {
		return last_id, errors.New(strconv.Itoa(int(id)) + "not found")
	}

}
func (repo *last_IdRepo) FindAll(ctx context.Context) ([]entity.Last_Id, error) {
	script := "SELECT id, email, text FROM last_id"
	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var last_id []entity.Last_Id
	for rows.Next() {
		last_id1 := entity.Last_Id{}
		rows.Scan(&last_id1.Id, &last_id1.Email, &last_id1.Text)
		last_id = append(last_id, last_id1)
	}
	return last_id, nil
}
