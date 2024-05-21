package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"base.url/class/fbufstat"
	"base.url/class/fwrite"
	"base.url/class/simplelogger"
	"github.com/gin-gonic/gin"
)

/*
type fstat struct {
	Id        int    `json:"id"`
	Timestamp string `json:"timestamp"`
}

*/

type gw struct {
	H_data []float64 `json:"h"`
	T_data []float64 `json:"t"`
}

const baseadm = "adm/"
const baseadmn = "StatDesc.json"

const basedata = "data/"

const baselog = "log/"
const baselogn = "LogStream.json"

func main() {

	StatDesc := fbufstat.NewObj(fwrite.UnFtoStrm(baseadm, baseadmn))

	router := gin.Default()

	router.GET("/stat", getStat(&StatDesc))
	router.GET("/dumpLogF", getLogF())

	router.POST("/sendF", postFile(&StatDesc))

	router.Run("localhost:8081")

}

func getStat(Table *fbufstat.Bufstat) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		ctx.IndentedJSON(http.StatusOK, Table)

	}

}

func getLogF() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var BufSl []simplelogger.LogWrite
		var BufLog simplelogger.LogWrite
		var bufst []string

		var count int

		bufst, count = fwrite.RFbyLine(baselog, baselogn)

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

		fgwsize, _ = fwrite.PrintStToF(basedata, fgwname, 0666, Gwbuf)

		simplelogger.LogWriteFile(baselog, baselogn, count, fgwsize, fgwname)

		//Updating app status
		Stpt.UpdateCnt()
		Stpt.UpdateSize(fgwsize)

		buf = Stpt.GetJSONObj()
		fwrite.PrintStrmToF(baseadm, baseadmn, 0666, buf)

		//Response
		ctx.IndentedJSON(http.StatusCreated, "Written "+string(fgwname)+"for "+strconv.Itoa(fgwsize)+" bytes")

	}

}
