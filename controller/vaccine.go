package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaaclyra132/petshop-go/config"
	"github.com/isaaclyra132/petshop-go/models"
)

func GetVaccines(c *gin.Context) {
	vaccines := []models.Vaccine{}
	config.DB.Find(&vaccines)
	c.JSON(200, &vaccines)
}

func CreateVaccine(c *gin.Context) {
	petID := c.Param("id")
	var pet models.Pet

	if err := config.DB.First(&pet, petID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Animal de estimação não encontrado"})
		return
	}

	var vacina models.Vaccine
	if err := c.ShouldBindJSON(&vacina); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vacina.PetID = pet.ID
	config.DB.Create(&vacina)

	c.JSON(http.StatusCreated, &vacina)
}

func UpdateVaccine(c *gin.Context) {
	var vaccine models.Vaccine
	config.DB.Where("id = ?", c.Param("id")).First(&vaccine)
	c.BindJSON(&vaccine)
	config.DB.Save(&vaccine)
	c.JSON(200, &vaccine)
}

func DeleteVaccine(c *gin.Context) {
	var vaccine models.Vaccine
	config.DB.Where("id = ?", c.Param("id")).Delete(&vaccine)
	c.JSON(204, "")
}
