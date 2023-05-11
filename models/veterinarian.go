package models

import (
	"time"

	"gorm.io/gorm"
)

type Veterinarian struct {
	gorm.Model
	ID                  uint      `json:"id" gorm:"primary_key"`
	Nome                string    `json:"nome"`
	DataNascimento      time.Time `json:"dataNascimento"`
	Certificacao        string    `json:"certificado"`
	VacinasAplicadas    []Vaccine
	ConsultasRealizadas []Appointment
}
