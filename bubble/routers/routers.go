package routers

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	//route
	r := gin.Default()
	//load static
	r.Static("/static", "./static")
	//load template
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHander)

	v1Group := r.Group("/v1")
	{
		//show all
		v1Group.GET("/todo", controller.ShowAll)
		//add one
		v1Group.POST("/todo", controller.AddOne)
		//update one
		v1Group.PUT("/todo/:id", controller.UpdateOne)
		//delete one
		v1Group.DELETE("/todo/:id", controller.DeleteOne)
	}

	return r
}
