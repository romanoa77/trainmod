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

	wrtch := make(chan fbufstat.Bufstat)

	if initstat(envdef.Baseadm, envdef.Baseadmn, StatDesc) != nil {

		simplelogger.LogPanic("FATAL ERROR", "FS ERROR")

	}

	simplelogger.LogGreet("DB ready")

	App := appstr.ConcrH{}

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	initapp(router, App, StatDesc, DSDesc, wrtch)

	srv := &http.Server{Addr: envdef.Basesrvurl, Handler: router}

	go worker(wrtch)

	srv.ListenAndServe()

	close(wrtch)

}

func worker(Bch chan fbufstat.Bufstat) {

	var Buf fbufstat.Bufstat
	var err error

	for {

		select {

		case Buf = <-Bch:

			_, err = fwrite.PrintStToF(envdef.Baseadm, envdef.Baseadmn, 0666, Buf)

			if err != nil {

				simplelogger.LogPanic("FATAL ERROR", "FS ERROR")

			}

		}

	}

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

func PostWho(Desc *dsstat.DSstat) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var buf dsstat.Idt

		ctx.BindJSON(&buf)

		Desc.User = buf.User
		Desc.Token = buf.Token

		ctx.Next()

	}
}

func initapp(SrvPt *gin.Engine, AppRts appmodel.AbstrApp, Buf *fbufstat.Bufstat, Desc *dsstat.DSstat, Chnl chan fbufstat.Bufstat) {

	SrvPt.GET("/stat", AppRts.GetStat(Buf))
	SrvPt.GET("/dumpLogF", AppRts.GetLogF())
	SrvPt.GET("/dstat", AppRts.GetDSdsc(Desc))
	SrvPt.POST("/sendF", AppRts.PostFile(Buf, Chnl))
	SrvPt.POST("/upddsc", PostWho(Desc), AppRts.PostDsc(Desc))
	SrvPt.POST("/cleanall", AppRts.PostClean(Buf, Desc))

}
