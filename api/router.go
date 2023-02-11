package api

import (
	"github.com/gin-gonic/gin"
	"todoList/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")
	todo := v1.Group("/todo")

	user.POST("/register", Register)
	user.POST("/login", Login)

	todo.Use(middleware.Auth())
	todo.POST("/create", C)
	todo.GET("/read", R)
	todo.PUT("/update", U)
	todo.DELETE("/delete", D)

	return r
}
