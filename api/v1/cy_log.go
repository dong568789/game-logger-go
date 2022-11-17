package v1

import (
	"game-logger-go/api/v1/msgdata"
	"game-logger-go/utils"
	"game-logger-go/utils/logSave"
	"github.com/gin-gonic/gin"
)

func LogWrite(c *gin.Context) {
	var req msgdata.ReqWriteLog
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.FailedResult(c, err.Error())
		return
	}
	var models []msgdata.ReqModel
	if err := c.ShouldBindJSON(&models); err != nil {
		utils.FailedResult(c, err.Error())
		return
	}
	params := make(map[string]string)
	params["PlayerID"] = req.PlayerID
	params["Date"] = req.Date
	params["pf"] = "yxgames"
	sign := utils.MakeSign(params)
	if sign != req.Token {
		utils.FailedResult(c, "sing error")
		return
	}
	for _, item := range models {
		if len(item.List) > 0 {
			for _, logItem := range item.List {
				logSave.CyMgr.WriteLog(item.Model, logItem)
			}
		}
	}
	utils.SuccessResult(c, "success")
}
