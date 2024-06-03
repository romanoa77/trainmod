package appstr

import (
	"base.url/class/fbufstat"
	"github.com/gin-gonic/gin"
)

type ConcrH struct{}

func (h ConcrH) GetStat(Table *fbufstat.Bufstat) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		ctx.JSON(200, Table)

	}

}
