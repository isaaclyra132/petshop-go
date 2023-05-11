package models

import "gorm.io/gorm"

type Appointment struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primary_key"`
	Sintomas    string `json:"sintomas"`
	Diagnostico string `json:"diagnostico"`
	Receita     string `json:"receitaMedica"`
	Informacoes string `json:"informacoesAdicionais"`
	PetID       uint
	// Pet            Pet
	VeterinarianID uint
	// Veterinarian   Veterinarian `json:"veterinario"`
}
