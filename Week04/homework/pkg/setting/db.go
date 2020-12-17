package setting

import (
	"database/sql"
	"fmt"
)

func NewDB(s *Setting) (*sql.DB, error) {
	section := &MysqlConfig{}
	if err := s.vp.UnmarshalKey("mysql", section); err != nil {
		return nil, err
	}
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%d@%s:%s/%s", section.host, section.port, section.user, section.password, section.db))
	return db, nil
}
