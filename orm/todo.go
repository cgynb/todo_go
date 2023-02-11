package orm

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func (t *Todo) AfterCreate(tx *gorm.DB) (err error) {
	user, _ := GetUser("id", strconv.Itoa(int(t.UserId)))
	user.TodoNum++
	DB.Save(user)
	return
}
func (t *Todo) AfterDelete(tx *gorm.DB) (err error) {
	user, _ := GetUser("id", strconv.Itoa(int(t.UserId)))
	fmt.Println(user)
	user.TodoNum--
	DB.Save(user)
	return
}

func CreateTodo(userId uint, content string, endTime time.Time) (*Todo, string, bool) {
	todo := &Todo{
		UserId:    userId,
		Status:    TODO,
		Content:   content,
		StartTime: time.Now(),
		EndTime:   endTime,
	}
	result := DB.Create(todo)
	if result.Error != nil {
		return nil, "please create todo again", false
	}
	return todo, "ok", true
}
func GetTodos(userId uint, page int) ([]*Todo, bool) {
	var Todos []*Todo
	result := DB.Where("user_id = ?", userId).Limit(pageSize).Offset(pageSize * (page - 1)).Find(&Todos)
	return Todos, result.Error == nil
}

func GetTodosByKey(userId uint, keyWord string, page int) ([]*Todo, bool) {
	var Todos []*Todo
	result := DB.Where("user_id = ? AND content LIKE ?", userId, "%"+keyWord+"%").Limit(pageSize).Offset(pageSize * (page - 1)).Find(&Todos)
	return Todos, result.Error == nil
}

func UpdateOneTodoStatus(userId, todoId uint, status string) bool {
	result := DB.Model(&Todo{UserId: userId, Model: gorm.Model{ID: todoId}}).Update("status", status)
	return result.Error == nil
}
func UpdateAllTodoStatus(userId uint, status string) bool {
	result := DB.Model(&Todo{}).Where("user_id = ?", userId).Update("status", status)
	return result.Error == nil
}
func DelTodo(userId, todoId uint) bool {
	result := DB.Where("user_id = ? AND id = ?", userId, todoId).Delete(&Todo{UserId: userId})
	return result.Error == nil
}

func DelTodosByStatus(userId uint, status string) bool {
	result := DB.Where("user_id = ? AND status = ?", userId, status).Delete(&Todo{UserId: userId})
	return result.Error == nil
}

func DelTodos(userId uint) bool {
	result := DB.Where("user_id = ?", userId).Delete(&Todo{UserId: userId})
	return result.Error == nil
}
