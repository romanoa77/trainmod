package appmodel

import (
	"base.url/class/dsstat"
	"base.url/class/fbufstat"
	"github.com/gin-gonic/gin"
)

type AbstrApp interface {
	GetStat(Table *fbufstat.Bufstat) gin.HandlerFunc
	GetLogF() gin.HandlerFunc
	PostFile(Stpt *fbufstat.Bufstat, Bch chan fbufstat.Bufstat) gin.HandlerFunc
	GetDSdsc(Table *dsstat.DSstat) gin.HandlerFunc
	PostDsc(Table *dsstat.DSstat) gin.HandlerFunc
	PostClean(Stpt *fbufstat.Bufstat, Table *dsstat.DSstat) gin.HandlerFunc
}
