package models

import "gorm.io/gorm"

type Categoria struct {
	gorm.Model
	Nombre  string
	TareaID uint
}
