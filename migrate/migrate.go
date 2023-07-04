package main

import (
	"github.com/Ushio-dev/tareas-app/initializers"
	"github.com/Ushio-dev/tareas-app/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Tarea{})
	initializers.DB.AutoMigrate(&models.Categoria{})
}
