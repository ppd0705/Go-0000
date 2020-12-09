package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"homework/errorcode"
	"homework/service"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/db?charset=utf8&timeout=1s&readTimeout=1s")
	if err != nil {
		panic(fmt.Sprintf("mysql open err: %v", err))
	}
}

func main() {
	obj, err := service.GetData(context.Background(), db, 1024)
	if errors.Is(err, errorcode.ErrNotFound) {
		log.Printf("404, %v", err)
		return
	}
	if err != nil {
		log.Printf("500, %v", err)
	}
	log.Printf("obj: %v", obj)
}
