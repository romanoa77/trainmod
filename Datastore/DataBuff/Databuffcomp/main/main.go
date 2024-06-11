package main

import (
	"net/http"

	"base.url/class/appmodel"
	"base.url/class/appstr"
	"base.url/class/envdef"
	"base.url/class/fbufstat"
	"base.url/class/fwrite"
	"base.url/class/simplelogger"
	"github.com/gin-gonic/gin"
)

func main() {

	StatDesc := fbufstat.New(0, 0)

	if initstat(envdef.Baseadm, envdef.Baseadmn, &StatDesc) != nil {

		simplelogger.LogPanic("FATAL ERROR", "FS ERROR")
	}

	simplelogger.LogGreet("DB ready")

	App := appstr.ConcrH{}

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	initapp(router, App, &StatDesc)

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

func initapp(SrvPt *gin.Engine, AppRts appmodel.AbstrApp, Buf *fbufstat.Bufstat) {

	SrvPt.GET("/stat", AppRts.GetStat(Buf))
	SrvPt.GET("/dumpLogF", AppRts.GetLogF())
	SrvPt.POST("/sendF", AppRts.PostFile(Buf))

}
