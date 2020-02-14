package app

import (
	"github.com/gin-gonic/gin"

	"capulus/package/w"
)

type Gin struct {
	C *gin.Context
}

// JSON alias type
type JSON = map[string]interface{}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	Status := "success"
	if httpCode != 200 {
		Status = "error"
	}
	g.C.JSON(httpCode, Response{
		Code:   errCode,
		Status: Status,
		Msg:    w.GetMsg(errCode),
		Data:   data,
	})
	return
}
