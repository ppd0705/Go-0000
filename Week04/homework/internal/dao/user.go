package dao

import (
	"database/sql"
	"homework/internal/model"
)

type Dao struct {
	db *sql.DB
}

func New(db *sql.DB) *Dao {
	return &Dao{db: db}
}

func (d *Dao) GetUser(id int) (*model.User, error) {
	u := &model.User{}
	row := d.db.QueryRow("select id, name from user where id = ?", id)
	err := row.Scan(&u.ID, &u.Name)
	if err != nil {
		return nil, err
	}
	return u, nil
}
