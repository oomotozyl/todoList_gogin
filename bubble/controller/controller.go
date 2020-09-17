package controller

import (
	"net/http"

	"bubble/models"
	"bubble/service"

	"github.com/gin-gonic/gin"
)

// IndexHander IndexHander
func IndexHander(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// ShowAll ShowAll
func ShowAll(c *gin.Context) {
	todoList, err := service.SelectAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

// AddOne AddOne
func AddOne(c *gin.Context) {
	var todo models.Todo //??? models　该不该放到service层去 把c直接传过去???
	//request
	c.BindJSON(&todo)
	//then response
	if err := service.CreateOne(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// UpdateOne UpdateOne
func UpdateOne(c *gin.Context) {
	//request
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "param id error"})
		return
	}
	//DB serch check
	todo, err := service.SelectOne(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	//DB insert
	c.BindJSON(todo)
	if err := service.UpdateOne(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"id": id})
	}

}

// DeleteOne DeleteOne
func DeleteOne(c *gin.Context) {
	//request
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "param id error"})
		return
	}
	if err := service.DeleteOne(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"}) //??? why return id:deleted ???
	}
}
