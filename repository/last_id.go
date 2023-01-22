package repository

import (
	"context"
	"go-database/entity"
)

type Last_IdRepository interface {
	Insert(ctx context.Context, last_Id entity.Last_Id) (entity.Last_Id, error)
	FindById(ctx context.Context, id int32) (entity.Last_Id, error)
	FindAll(ctx context.Context) ([]entity.Last_Id, error)
}
