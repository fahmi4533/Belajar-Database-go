package repository

import (
	"context"
	"fmt"
	godatabase "go-database"
	"go-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsert(t *testing.T) {
	last_idRepository := NewRepository(godatabase.GetConnection())

	ctx := context.Background()
	last_id := entity.Last_Id{
		Email: "fahmi@gma.com",
		Text:  "jamaaaah",
	}
	result, err := last_idRepository.Insert(ctx, last_id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindId(t *testing.T) {
	last_idRepository2 := NewRepository(godatabase.GetConnection())

	// ctx := context.Background()
	cmnt, err := last_idRepository2.FindById(context.Background(), 17)
	if err != nil {
		panic(err)
	}
	fmt.Println(cmnt)
}
func TestFindAll(t *testing.T) {
	last_idRepository2 := NewRepository(godatabase.GetConnection())

	// ctx := context.Background()
	cmnt, err := last_idRepository2.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, v := range cmnt {
		fmt.Println(v)
	}
}
