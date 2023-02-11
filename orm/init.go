package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todoList/config"
)

func InitDB() {
	db, err := gorm.Open(mysql.Open(config.Conf.MysqlConfig.DNS), &gorm.Config{})
	DB = db
	err = DB.AutoMigrate(&User{}, &Todo{})
	if err != nil {
		panic("db err")
	}
}
