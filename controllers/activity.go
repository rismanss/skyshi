package controllers

import (
	"net/http"
	"skyshi/config"
	"skyshi/models"

	"github.com/gin-gonic/gin"
)

func GetAllActivity(c *gin.Context) {
	var activity []models.Activity

	config.DB.Find(&activity)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    activity,
	})
}

func CreateActivity(c *gin.Context) {
	var activity models.Activity

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&activity)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    activity,
	})
}

func GetByIdActivity(c *gin.Context) {
	var activity models.Activity
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&activity).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found !"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    activity,
	})
}

func UpdateActivity(c *gin.Context) {
	var activity models.Activity
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&activity).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found !"})
		return
	}

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&activity)
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    activity,
	})
}

func DeleteActivity(c *gin.Context) {
	var activity models.Activity
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&activity).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	config.DB.Delete(&activity)
	c.JSON(http.StatusOK, gin.H{
		"status":  "Deleted",
		"message": "Success",
	})
}
