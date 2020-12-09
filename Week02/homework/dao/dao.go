package dao

import (
	"context"
	"database/sql"
	"fmt"
	xerrors "github.com/pkg/errors"
	"homework/errorcode"
	"homework/model"
)

func GetObj(ctx context.Context, db *sql.DB, id int64) (*model.Article, error) {
	obj := &model.Article{}
	row := db.QueryRowContext(ctx, "select id, title, content from article where id = %?", id)
	err := row.Scan(&obj.ID, &obj.Title, &obj.Content)
	if err == sql.ErrNoRows {
		return nil, xerrors.Wrapf(errorcode.ErrNotFound, fmt.Sprintf("sql err: %s", err))
	}
	if err != nil {
		return nil, xerrors.Wrap(err, "Failed to get article")
	}
	return obj, nil
}
