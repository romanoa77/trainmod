package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"base.url/class/appmodel"
	"base.url/class/appstr"
	"base.url/class/envdef"
	"base.url/class/fbufstat"
	"base.url/class/fwrite"
	"base.url/class/simplelogger"
	"github.com/gin-gonic/gin"
)

type gw struct {
	H_data []float64 `json:"h"`
	T_data []float64 `json:"t"`
}

/*Note: Simple I/O functions can print messages*/

/*
export ADMROOT = "adm/"
const baseadmn = "StatDesc.json"

const basedata = "data/"

const baselog = "log/"
const baselogn = "LogStream.json"
*/

func main() {

	StatDesc := fbufstat.NewObj(fwrite.UnFtoStrm(envdef.Baseadm, envdef.Baseadmn))
	App := appstr.ConcrH{}

	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	initapp(router, App, &StatDesc)

	srv := &http.Server{Addr: ":8080", Handler: router}

	srv.ListenAndServe()

	//router.GET("/dumpLogF", getLogF())

	//router.POST("/sendF", postFile(&StatDesc))

	router.Run(envdef.Basesrvurl)

}

func initapp(SrvPt *gin.Engine, AppRts appmodel.AbstrApp, Buf *fbufstat.Bufstat) {

	SrvPt.GET("/stat", AppRts.GetStat(Buf))
	//router.GET(, getStat(&StatDesc))

}

func getLogF() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var BufSl []simplelogger.LogWrite
		var BufLog simplelogger.LogWrite
		var bufst []string

		var count int

		bufst, count = fwrite.RFbyLine(envdef.Baselog, envdef.Strmlogn)

		for i := 0; i < count; i++ {

			lineb := []byte(bufst[i])
			json.Unmarshal(lineb, &BufLog)
			BufSl = append(BufSl, BufLog)

		}

		ctx.IndentedJSON(http.StatusOK, BufSl)

	}

}

func postFile(Stpt *fbufstat.Bufstat) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var Gwbuf gw
		var count int
		var fgwsize int
		var fgwname string
		var buf []byte

		count = Stpt.N_itm

		fgwname = "strmgw" + strconv.Itoa(count) + ".json"

		// Call BindJSON to bind the received JSON to
		// newAlbum.

		if err := ctx.BindJSON(&Gwbuf); err != nil {
			return
		}

		fgwsize, _ = fwrite.PrintStToF(envdef.Basedata, fgwname, 0666, Gwbuf)

		simplelogger.LogWriteFile(envdef.Baselog, envdef.Strmlogn, count, fgwsize, fgwname)

		//Updating app status
		Stpt.UpdateCnt()
		Stpt.UpdateSize(fgwsize)

		buf = Stpt.GetJSONObj()
		fwrite.PrintStrmToF(envdef.Baseadm, envdef.Baseadmn, 0666, buf)

		//Response
		ctx.IndentedJSON(http.StatusCreated, "Written "+string(fgwname)+"for "+strconv.Itoa(fgwsize)+" bytes")

	}

}
