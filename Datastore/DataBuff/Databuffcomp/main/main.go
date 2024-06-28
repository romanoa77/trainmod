package main

import (
	"net/http"

	"base.url/class/appmodel"
	"base.url/class/appstr"
	"base.url/class/dsstat"
	"base.url/class/envdef"
	"base.url/class/fbufstat"
	"base.url/class/fwrite"
	"base.url/class/simplelogger"
	"github.com/gin-gonic/gin"
)

func main() {

	StatDesc := fbufstat.GetInst()
	DSDesc := dsstat.GetInst()

	if initstat(envdef.Baseadm, envdef.Baseadmn, StatDesc) != nil {

		simplelogger.LogPanic("FATAL ERROR", "FS ERROR")

	}

	simplelogger.LogGreet("DB ready")

	App := appstr.ConcrH{}

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	initapp(router, App, StatDesc, DSDesc)

	srv := &http.Server{Addr: envdef.Basesrvurl, Handler: router}

	srv.ListenAndServe()

}

func initstat(rtdir string, rtfname string, Buf *fbufstat.Bufstat) error {
	var err error
	var cbuf []byte

	cbuf, err = fwrite.UnFtoStrm(rtdir, rtfname)

	if err == nil {

		Buf.SetStat(cbuf)
	}

	return err
}

func PostMid(Stpt *fbufstat.Bufstat) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		ctx.Next()

		var Statbuf fbufstat.Bufstat
		var err error

		size := Stpt.Buff_size
		count := Stpt.N_itm

		Statbuf.N_itm = count
		Statbuf.Buff_size = size
		_, err = fwrite.AtmWrtJs(envdef.Baseadm, envdef.Baseadmn, Statbuf)

		if err != nil {

			simplelogger.LogPanic("FATAL ERROR", "FS ERROR")

		}

	}
}

func PostWho(Desc *dsstat.DSstat) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var buf dsstat.Idt

		ctx.BindJSON(&buf)

		Desc.User = buf.User
		Desc.Token = buf.Token

		ctx.Next()

	}
}

func initapp(SrvPt *gin.Engine, AppRts appmodel.AbstrApp, Buf *fbufstat.Bufstat, Desc *dsstat.DSstat) {

	SrvPt.GET("/stat", AppRts.GetStat(Buf))
	SrvPt.GET("/dumpLogF", AppRts.GetLogF())
	SrvPt.GET("/dstat", AppRts.GetDSdsc(Desc))
	SrvPt.POST("/sendF", PostMid(Buf), AppRts.PostFile(Buf))
	SrvPt.POST("/upddsc", PostWho(Desc), AppRts.PostDsc(Desc))
	SrvPt.POST("/cleanall", AppRts.PostClean(Buf, Desc))

}
