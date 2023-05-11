package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaaclyra132/petshop-go/controller"
)

func AppRoutes(router *gin.Engine) {
	// Rotas User
	router.GET("/users", controller.GetUsers)
	router.POST("/users", controller.CreateUser)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)

	// Rotas Pet
	router.GET("/pets", controller.GetPets)
	router.POST("/users/:id/pets", controller.CreatePet)
	router.PUT("/pets/:id", controller.UpdatePet)
	router.DELETE("/pets/:id", controller.DeletePet)

	// Rotas Vaccine
	router.GET("/vaccines", controller.GetVaccines)
	router.POST("/pets/:id/vaccines", controller.CreateVaccine)
	router.PUT("/vaccines/:id", controller.UpdateVaccine)
	router.DELETE("/vaccines/:id", controller.DeleteVaccine)

	// Rotas Appointment
	router.GET("/appointments", controller.GetAppointments)
	router.POST("/pets/:id/appointments", controller.CreateAppointment)
	router.PUT("/appointments/:id", controller.UpdateAppointment)
	router.DELETE("/appointments/:id", controller.DeleteAppointment)

	// Rotas Veterinarian
	router.GET("/veterinarians", controller.GetVeterinarians)
	router.POST("/veterinarians/:id", controller.CreateVeterinarian)
	router.PUT("/veterinarians/:id", controller.UpdateVeterinarian)
	router.DELETE("/veterinarians/:id", controller.DeleteVeterinarian)
}
