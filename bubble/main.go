/*
get /         >> index.html
get /v1/todo  >> 200,{todoList:todoList}
post /v1/todo >> 200,{nil}
put  /v1/todo/:id   >> 200,{nil}
delete /v1/todo/:id >> 200,{nil}
*/
package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {
	//DB connect
	err := dao.InitPostgs()
	if err != nil {
		panic(err)
	}
	//table init
	dao.DB.AutoMigrate(models.Todo{})

	//router set
	r := routers.SetRouter()

	//run server 8080
	r.Run()
}
