package appmodel

import (
	"base.url/class/fbufstat"
	"github.com/gin-gonic/gin"
)

type AbstrApp interface {
	GetStat(Table *fbufstat.Bufstat) gin.HandlerFunc
}
