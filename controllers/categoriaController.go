package controllers

import (
	"github.com/Ushio-dev/tareas-app/initializers"
	"github.com/Ushio-dev/tareas-app/models"
	"github.com/gin-gonic/gin"
)

func GetCategorias(c *gin.Context) {
	var categorias []models.Categoria
	initializers.DB.Find(&categorias)

	c.JSON(200, gin.H{
		"categorias": categorias,
	})
}
