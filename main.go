package main

import (
	"todoList/api"
	"todoList/config"
	"todoList/orm"
)

func main() {
	orm.InitDB()
	config.InitConf()
	api.Router().Run(":80")
}
