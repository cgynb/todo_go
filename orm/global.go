package orm

import (
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

const (
	TODO  = 1
	DOING = 2
	DONE  = 3
)

const pageSize = 5

type User struct {
	Name     string `json:"name" gorm:"unique"`
	Password string `json:"-"`
	Email    string `json:"email"`
	TodoNum  uint   `json:"todo_num"`
	gorm.Model
}
type Todo struct {
	UserId    uint      `json:"user_id"`
	Status    int8      `json:"status" gorm:"default:1"`
	Content   string    `json:"content"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	gorm.Model
}
