package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaaclyra132/petshop-go/config"
	"github.com/isaaclyra132/petshop-go/models"
)

func GetAppointments(c *gin.Context) {
	appointments := []models.Appointment{}
	config.DB.Find(&appointments)
	c.JSON(200, &appointments)
}

func CreateAppointment(c *gin.Context) {
	petID := c.Param("id")
	var pet models.Pet

	if err := config.DB.First(&pet, petID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Animal de estimação não encontrado"})
		return
	}

	var consulta models.Appointment
	if err := c.ShouldBindJSON(&consulta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	consulta.PetID = pet.ID
	config.DB.Create(&consulta)

	c.JSON(http.StatusCreated, &consulta)
}

func UpdateAppointment(c *gin.Context) {
	var appointment models.Appointment
	config.DB.Where("id = ?", c.Param("id")).First(&appointment)
	c.BindJSON(&appointment)
	config.DB.Save(&appointment)
	c.JSON(200, &appointment)
}

func DeleteAppointment(c *gin.Context) {
	var appointment models.Appointment
	config.DB.Where("id = ?", c.Param("id")).Delete(&appointment)
	c.JSON(204, "")
}
