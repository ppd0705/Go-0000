package scripts

import (
	"github.com/google/wire"
	"homework/internal/biz"
	"homework/internal/dao"
	"homework/internal/server"
	"homework/pkg/setting"
)

func Initial(addr string) (*server.Server, error){
	wire.Build(setting.NewConfig,setting.NewDB, dao.New, biz.New, server.New)
	return &server.Server{}, nil
}
