package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"todoList/orm"
	"todoList/response"
)

func C(c *gin.Context) {
	var user *orm.User
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*orm.User)
	}
	content := c.PostForm("content")
	endTime := c.PostForm("end_time")
	et, _ := strconv.Atoi(endTime)
	orm.CreateTodo(user.ID, content, time.Unix(int64(et), 0))
	c.JSON(http.StatusOK, response.RespOk(nil))
}
func R(c *gin.Context) {
	var user *orm.User
	var todos []*orm.Todo
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*orm.User)
	}
	page, _ := strconv.Atoi(c.Query("page"))
	keyWord, ok := c.GetQuery("keyword")
	if ok {
		todos, _ = orm.GetTodosByKey(user.ID, keyWord, page)
	} else {
		todos, _ = orm.GetTodos(user.ID, page)
	}
	c.JSON(http.StatusOK, response.RespOk(gin.H{
		"todos": todos,
		"page":  page,
	}))
}
func U(c *gin.Context) {
	var user *orm.User
	var ok bool
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*orm.User)
	}
	num := c.PostForm("num")
	status := c.PostForm("status")

	switch num {
	case "one":
		todoId := c.PostForm("todo_id")
		t, _ := strconv.Atoi(todoId)
		ok = orm.UpdateOneTodoStatus(user.ID, uint(t), status)
	case "all":
		ok = orm.UpdateAllTodoStatus(user.ID, status)
	}

	if ok {
		c.JSON(http.StatusOK, response.RespOk(nil))
	} else {
		c.JSON(http.StatusOK, response.RespErr(http.StatusForbidden, "please update again", "update fail"))
	}
}
func D(c *gin.Context) {
	var user *orm.User
	var ok bool
	if val, ok := c.Get("user"); ok && val != nil {
		user, _ = val.(*orm.User)
	}
	num := c.PostForm("num")
	switch num {
	case "one":
		todoId := c.PostForm("todo_id")
		t, _ := strconv.Atoi(todoId)
		ok = orm.DelTodo(user.ID, uint(t))
	case "all":
		status, hasStatus := c.GetPostForm("status")
		if hasStatus {
			ok = orm.DelTodosByStatus(user.ID, status)
		} else {
			ok = orm.DelTodos(user.ID)
		}
	}
	if ok {
		c.JSON(http.StatusOK, response.RespOk(nil))
	} else {
		c.JSON(http.StatusOK, response.RespErr(http.StatusForbidden, "please del again", "del fail"))
	}
}
