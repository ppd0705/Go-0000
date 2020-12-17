package main

import (
	"flag"
	"homework/internal/biz"
	"homework/internal/dao"
	"homework/internal/server"
	"homework/pkg/setting"
	"log"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":0", "")
	flag.Parse()
}

func GetServer() (*server.Server, error) {
	settingSetting, err := setting.NewConfig()
	if err != nil {
		return nil, err
	}
	db, err := setting.NewDB(settingSetting)
	if err != nil {
		return nil, err
	}
	daoDao := dao.New(db)
	bizBiz := biz.New(daoDao)
	serverServer := server.New(addr, bizBiz)
	return serverServer, nil
}
func main() {
	s, err := GetServer()
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	s.Start()
}
