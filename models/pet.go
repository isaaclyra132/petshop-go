package models

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	ID             uint      `json:"id" gorm:"primary_key"`
	Nome           string    `json:"nome"`
	Peso           float64   `json:"peso"`
	Raca           string    `json:"raca"`
	Especie        string    `json:"especie"`
	Sexo           string    `json:"sexo"`
	DataNascimento time.Time `json:"dataNascimento"`
	UserID         uint      `json:"tutorId"`
	// User           User          `json:"tutor"`
	Consultas []Appointment `json:"consultas"`
	Vacinas   []Vaccine     `json:"vacinas"`
}
