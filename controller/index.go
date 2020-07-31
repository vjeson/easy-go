package controller

import (
	"demo/middleware"
	"demo/service"
	"demo/util"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	//读取body1
	//buf := make([]byte, 1024)
	//n, _ := c.Request.Body.Read(buf)
	//body := string(buf[0:n])
	//读取body2
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Print(data)
	var user service.User
	err := c.BindJSON(&user)
	if err != nil {
		res := &util.Response{Code: 1000, Msg: "数据解析出错"}
		res.Json(c)
		return
	}
	u := user.ChekcLogin()
	if u == nil {
		res := &util.Response{Code: 1000, Msg: "用户名或密码错误"}
		res.Json(c)
		return
	}
	jwt := middleware.NewJWT()
	token, _ := jwt.GenerateToken(u)
	res := &util.Response{Code: 0, Msg: "ok", Data: token}
	res.Json(c)
}

func Hello(c *gin.Context) {
	data := map[string]interface{}{"abc":123}
	res := &util.Response{
		Code: 0,
		Msg:  "",
		Data: data,
	}

	res.Json(c)
}
