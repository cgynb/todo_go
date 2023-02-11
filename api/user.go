package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"todoList/middleware"
	"todoList/orm"
	"todoList/response"
	"todoList/utils"
)

func Register(c *gin.Context) {
	name := c.PostForm("username")
	email := c.PostForm("email")
	pwd := c.PostForm("password")
	user, ok, msg := orm.CreateUser(name, email, pwd)
	if ok {
		c.JSON(http.StatusOK, response.RespOk(user))
	} else {
		c.JSON(http.StatusForbidden, response.RespErr(403, msg, "register error"))
	}
}

func Login(c *gin.Context) {
	name := c.PostForm("username")
	pwd := c.PostForm("password")
	user, ok := orm.GetUser("name", name)
	if ok && utils.CheckPassword(pwd, user.Password) {
		token, _ := middleware.GenToken(&middleware.UserClaims{
			User:           *user,
			StandardClaims: jwt.StandardClaims{},
		})
		c.Header("Authorization", token)
		c.JSON(http.StatusOK, response.RespOk(user))
	} else {
		c.JSON(http.StatusForbidden, response.RespErr(403, "please login again", "name doesn't exist or password error"))
	}
}
