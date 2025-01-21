package controller

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/dto"
	"superMarket-backend/response"
	"superMarket-backend/service"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := &dto.User{
		Username: username,
		Password: password,
	}
	if user == nil {
		response.Fail(c, nil, "用户名和密码不能为空")
	}
	service.UserLogin(c, user)
}

func CheckedToken(c *gin.Context) {
	token := c.Query("token")
	service.CheckedToken(c, token)
}

func EmpMenu(c *gin.Context) {
	token := c.GetHeader("token")
	service.EmpMenu(c, token)
}

func Exit(c *gin.Context) {
	token := c.GetHeader("token")
	service.Exit(c, token)

}

func Logout(c *gin.Context) {
	content := c.PostForm("content")
	token := c.GetHeader("token")
	service.Logout(c, token, content)
}
