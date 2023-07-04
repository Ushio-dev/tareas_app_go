package models

import "gorm.io/gorm"

type Tarea struct {
	gorm.Model
	Nombre    string
	Categoria Categoria
}
