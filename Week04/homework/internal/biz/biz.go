package biz

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"homework/internal/dao"
	"strconv"
)

type Biz struct {
	dao *dao.Dao
}

func New(dao *dao.Dao) *Biz {
	return &Biz{dao: dao}
}

func (b *Biz) GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.GetString("id"))
	user, _ := b.dao.GetUser(id)
	fmt.Printf("user: %v", user)
}
