package controllers

import (
	"net/http"
	"skyshi/config"
	"skyshi/models"

	"github.com/gin-gonic/gin"
)

func GetAllTodo(c *gin.Context) {
	var todo []models.Todo

	config.DB.Find(&todo)
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})
}

func GetByIdTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "Not Found !", "message": "id " + id + " not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	todo.Priority = "very-high"

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&todo)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "Not Found !", "message": "id " + id + " not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&todo)
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "Not Found !", "message": "id " + id + " not found"})
		return
	}

	config.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{
		"status":  "Deleted",
		"message": "Success",
	})
}
