package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code 	int         `json:"code"`
	Msg  	string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

func (res *Response) Json(c *gin.Context)  {
	c.JSON(http.StatusOK, res)
	return
}
