/*
get /   >> index.html
get    /v1/todo >> 200,{todoList:todoList}
post   /v1/todo >> 200,{nil}
put    /v1/todo/:id >> 200,{nil}
delete /v1/todo/:id >> 200,{nil}

*/
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func InitPostgs() (err error) {
	dsn := "host=127.0.0.1 user=postgres password=password dbname=gogin port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) //???DB global,err local, how??? 解决：两个返回值时，可以提出返回err的方法
	if err != nil {
		return err
	}
	return
}

func main() {
	//DB connect
	err := InitPostgs()
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(Todo{})

	//route
	r := gin.Default()
	//load static
	r.Static("/static", "./static")
	//load template
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("/v1")
	{
		//show all
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []*Todo
			if err := DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		//add one
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			//request
			c.BindJSON(&todo)
			//then response
			if err := DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//update one
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			//request
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "param id error"})
				return
			}
			//DB serch check
			var todo Todo
			if err := DB.First(&todo, id).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			//DB insert
			c.BindJSON(&todo)
			if err := DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"id": id})
			}

		})
		//delete one
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			//request
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "param id error"})
				return
			}
			if err := DB.Delete(&Todo{}, id).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"}) //??? why return id:deleted ???
			}
		})
	}
	//run
	r.Run()
}
