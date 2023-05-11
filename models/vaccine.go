package models

import (
	"time"

	"gorm.io/gorm"
)

type Vaccine struct {
	gorm.Model
	ID            uint      `json:"id" gorm:"primary_key"`
	Nome          string    `json:"nome" gorm:"primary_key"`
	TipoVacina    string    `json:"tipoVacina"`
	DataAplicacao time.Time `json:"dataAplicacao"`
	PetID         uint
	// Pet            Pet `json:"pet"`
	VeterinarianID uint
	// Veterinarian   Veterinarian `json:"veterinario"`
}
