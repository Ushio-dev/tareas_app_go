package main

import (
	"github.com/Ushio-dev/tareas-app/controllers"
	"github.com/Ushio-dev/tareas-app/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/", controllers.TareaCreate)
	r.GET("/", controllers.GetTareas)
	r.GET("/categorias", controllers.GetCategorias)
	r.GET("/:id", controllers.GetTarea)
	r.Run()
}
