package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/isaaclyra132/petshop-go/config"
	"github.com/isaaclyra132/petshop-go/models"
)

func GetVeterinarians(c *gin.Context) {
	veterinarians := []models.User{}
	config.DB.Find(&veterinarians)
	c.JSON(200, &veterinarians)
}

func CreateVeterinarian(c *gin.Context) {
	var veterinarian models.Veterinarian
	c.BindJSON(&veterinarian)
	config.DB.Create(&veterinarian)
	c.JSON(200, &veterinarian)
}

func UpdateVeterinarian(c *gin.Context) {
	var veterinarian models.Veterinarian
	config.DB.Where("id = ?", c.Param("id")).First(&veterinarian)
	c.BindJSON(&veterinarian)
	config.DB.Save(&veterinarian)
	c.JSON(200, &veterinarian)
}

func DeleteVeterinarian(c *gin.Context) {
	var veterinarian models.Veterinarian
	config.DB.Where("id = ?", c.Param("id")).Delete(&veterinarian)
	c.JSON(204, "")
}
