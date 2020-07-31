package controller

import (
	"demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Users(c *gin.Context)  {
	var user service.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	result, err := user.Users()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1000,
			"msg": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": "",
		"data":   result,
	})
}

//添加数据
func Store(c *gin.Context) {
	var user service.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	id, err := user.Insert()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1000,
			"msg": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg": "添加成功",
		"data":    id,
	})
}

//修改数据
func Update(c *gin.Context) {
	var user service.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.Update(id)
	if err != nil || result.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    1000,
			"msg": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg": "修改成功",
	})
}

//删除数据
func Destroy(c *gin.Context) {
	var user service.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Destroy(id)
	if err != nil || result.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "删除成功",
	})
}
