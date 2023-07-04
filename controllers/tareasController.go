package controllers

import (
	"github.com/Ushio-dev/tareas-app/initializers"
	"github.com/Ushio-dev/tareas-app/models"
	"github.com/gin-gonic/gin"
)

func TareaCreate(c *gin.Context) {
	/*
		var nuevaTarea struct {
			Nombre    string
			Categoria models.Categoria
		}

		c.Bind(&nuevaTarea)

	*/
	// de esta forma solo necesito el id de la categoria y no todo el json
	var nuevaTarea struct {
		Nombre    string
		Categoria models.Categoria
	}
	var newTarea struct {
		Nombre      string
		CategoriaId uint
	}

	c.Bind(&newTarea)
	var categoria models.Categoria
	initializers.DB.First(&categoria, newTarea.CategoriaId)
	nuevaTarea.Nombre = newTarea.Nombre
	nuevaTarea.Categoria = categoria

	tarea := models.Tarea{Nombre: nuevaTarea.Nombre, Categoria: nuevaTarea.Categoria}
	result := initializers.DB.Create(&tarea)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"msg": "Tarea creada",
	})
}

func GetTareas(c *gin.Context) {
	var tareas []models.Tarea
	initializers.DB.Model(&models.Tarea{}).Preload("Categoria").Find(&tareas)

	c.JSON(200, gin.H{
		"tareas": tareas,
	})
}

func GetTarea(c *gin.Context) {
	id := c.Param("id")
	var tarea models.Tarea

	initializers.DB.Model(&models.Tarea{}).Preload("Categoria").First(&tarea, id)

	c.JSON(200, gin.H{
		"tarea": tarea,
	})
}
