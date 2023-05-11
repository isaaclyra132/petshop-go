package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaaclyra132/petshop-go/config"
	"github.com/isaaclyra132/petshop-go/models"
)

func GetPets(c *gin.Context) {
	pets := []models.Pet{}
	config.DB.Find(&pets)
	c.JSON(200, &pets)
}

func CreatePet(c *gin.Context) {
	userID := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	var pet models.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet.UserID = user.ID
	config.DB.Create(&pet)

	c.JSON(http.StatusCreated, &pet)
}

func UpdatePet(c *gin.Context) {
	var pet models.Pet
	config.DB.Where("id = ?", c.Param("id")).First(&pet)
	c.BindJSON(&pet)
	config.DB.Save(&pet)
	c.JSON(200, &pet)
}

func DeletePet(c *gin.Context) {
	var pet models.Pet
	config.DB.Where("id = ?", c.Param("id")).Delete(&pet)
	c.JSON(204, "")
}
