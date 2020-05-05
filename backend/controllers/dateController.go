package controller

import (
	"net/http"

	"github.com/aloksingh3112/BookMyMovie/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetDate(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var date models.Date
	err := db.Where("id=?", c.Param("id")).First(&date).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Something went wrong", "statusCode": 500})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": date, "message": "Fetch Data Successfully", "statusCode": 200})
}
