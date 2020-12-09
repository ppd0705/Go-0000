package service

import (
	"context"
	"database/sql"
	"homework/dao"
	"homework/model"
)

func GetData(ctx context.Context, db *sql.DB, id int64) (*model.Article, error) {
	return dao.GetObj(ctx, db, id)
}
