package appstr

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"base.url/class/dsstat"
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

func (h ConcrH) GetDSdsc(Table *dsstat.DSstat) gin.HandlerFunc {

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

func (h ConcrH) PostFile(Stpt *fbufstat.Bufstat, Bch chan fbufstat.Bufstat) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var Gwbuf gw

		var fgwsize int
		var fgwname string

		var err, errb error
		var resp_code int
		var resp any

		//Stpt.UpdateCnt()

		cTime := time.Now()
		fgwname = "strmgw" + cTime.Format("15:04:05.000000000") + ".json"

		// Call BindJSON to bind the received JSON to
		// newAlbum.

		errb = ctx.BindJSON(&Gwbuf)

		fgwsize, err = fwrite.AtmWrtJs(envdef.Basedata, fgwname, Gwbuf)

		if (err == nil) && (errb == nil) {

			Stpt.UpdateSt(fgwsize)
			Bch <- *Stpt

			simplelogger.LogWriteFile(envdef.Baselog, envdef.Strmlogn, Stpt.N_itm, fgwsize, fgwname)

			//Updating app status
			//Stpt.UpdateCnt()

			resp_code = http.StatusCreated
			resp = neterr.New("RES_CREATE", "Written "+string(fgwname)+"for "+strconv.Itoa(fgwsize)+" bytes")

		} else {
			resp_code = http.StatusInternalServerError
			resp = neterr.New(neterr.CdErrIO, neterr.BodyFIO)
			//Stpt.CancCnt()

		}

		ctx.JSON(resp_code, resp)

	}

}

func (h ConcrH) PostDsc(Table *dsstat.DSstat) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		Table.SetToFrz()

		resp_code := http.StatusCreated
		resp := neterr.New("STAT_UPD", "Table updated")

		ctx.JSON(resp_code, resp)

	}
}

func (h ConcrH) PostClean(Stpt *fbufstat.Bufstat, Table *dsstat.DSstat) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var buf fbufstat.Bufstat
		var flag bool
		var err error

		var resp_code int
		var resp any

		flag = false

		now := time.Now()

		if os.Rename(envdef.Basedata, envdef.Basevoldt+now.Format("Mon Jan 2 15:04")) != nil {

			//simplelogger.LogInf(envdef.Baselog, "infolog", "errore nominazione")
			flag = true
		}

		if os.Mkdir(envdef.Datanm, os.ModePerm) != nil {
			flag = true

			//simplelogger.LogInf(envdef.Baselog, "infolog", "errore creazione dir")

		}

		_, err = os.Create(envdef.Baselog + envdef.Strmlogn)

		if err != nil {

			//simplelogger.LogInf(envdef.Baselog, "infolog", "errore log")

			flag = true
		}

		Stpt.N_itm = 0
		Stpt.Buff_size = 0

		buf.N_itm = 0
		buf.Buff_size = 0

		Table.SetToOp()
		Table.User = "databuf"
		Table.Token = "databuf"

		_, err = fwrite.AtmWrtJs(envdef.Baseadm, envdef.Baseadmn, buf)

		if err != nil {

			//simplelogger.LogInf(envdef.Baselog, "infolog", "errore stato")

			flag = true
		}

		if flag {

			resp_code = http.StatusInternalServerError
			resp = neterr.New(neterr.CdErrIO, neterr.BodyFIO)

		} else {

			resp_code = http.StatusCreated
			resp = neterr.New("STAT_UPD", "Buffer cleaned")

		}

		ctx.JSON(resp_code, resp)

	}
}
