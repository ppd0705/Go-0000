package server

import (
	"github.com/gin-gonic/gin"
	"homework/internal/biz"
	"log"
	"net/http"
)

type Server struct {
	s *http.Server
}

func New(addr string, biz *biz.Biz) *Server {
	route := gin.New()
	route.GET("/user", biz.GetUser)
	s := &http.Server{
		Addr:    addr,
		Handler: route,
	}
	return &Server{s: s}

}

func (s *Server) Start() {
	if err := s.s.ListenAndServe(); err != nil {
		log.Fatalf("err: %v", err)
	}

}
