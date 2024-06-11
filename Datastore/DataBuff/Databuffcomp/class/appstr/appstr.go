package appstr

import (
	"encoding/json"
	"net/http"
	"strconv"

	"base.url/class/envdef"
	"base.url/class/fbufstat"
	"base.url/class/fwrite"
	"base.url/class/neterr"
	"base.url/class/simplelogger"
	"github.com/gin-gonic/gin"
)

type gw struct {
	H_data []float64 `json:"h"`
	T_data []float64 `json:"t"`
}
type ConcrH struct{}

func (h ConcrH) GetStat(Table *fbufstat.Bufstat) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, Table)

	}

}

func (h ConcrH) GetLogF() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var BufSl []simplelogger.LogWrite
		var BufLog simplelogger.LogWrite
		var bufst []string

		var count int
		var resp_code int
		var resp any

		bufst, count = fwrite.RFbyLine(envdef.Baselog, envdef.Strmlogn)

		if count != 0 {

			for i := 0; i < count; i++ {

				lineb := []byte(bufst[i])
				json.Unmarshal(lineb, &BufLog)
				BufSl = append(BufSl, BufLog)

			}

			resp_code = http.StatusOK
			resp = BufSl
		} else {
			resp_code = http.StatusNoContent
			resp = "{}"
		}

		ctx.JSON(resp_code, resp)

	}

}

func (h ConcrH) PostFile(Stpt *fbufstat.Bufstat) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var Gwbuf gw
		var count int
		var fgwsize int
		var fgwname string

		var err error
		var resp_code int
		var resp any

		count = Stpt.N_itm

		fgwname = "strmgw" + strconv.Itoa(count) + ".json"

		// Call BindJSON to bind the received JSON to
		// newAlbum.

		ctx.BindJSON(&Gwbuf)

		fgwsize, err = fwrite.AtmWrtJs(envdef.Basedata, fgwname, Gwbuf)

		if err == nil {

			simplelogger.LogWriteFile(envdef.Baselog, envdef.Strmlogn, count, fgwsize, fgwname)

			//Updating app status
			Stpt.UpdateCnt()
			Stpt.UpdateSize(fgwsize)

			_, err = fwrite.AtmWrtJs(envdef.Baseadm, envdef.Baseadmn, Stpt)

			if err == nil {
				resp_code = http.StatusCreated
				resp = neterr.New("RES_CREATE", "Written "+string(fgwname)+"for "+strconv.Itoa(fgwsize)+" bytes")

			} else {
				resp_code = http.StatusInternalServerError
				resp = neterr.New(neterr.CdErrIO, neterr.BodyFIO)

			}

		} else {
			resp_code = http.StatusInternalServerError
			resp = neterr.New(neterr.CdErrIO, neterr.BodyFIO)

		}

		ctx.JSON(resp_code, resp)

	}

}

//
