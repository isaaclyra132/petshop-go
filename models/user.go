package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	TipoUsuario string `json:"tipoUsuario"`
	Pets        []Pet
}
