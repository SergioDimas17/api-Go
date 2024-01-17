package main

import (
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"github.com/SergioDimas17/api-go/controllers"
=======
	"github.com/SergioDimas17/controllers"
>>>>>>> 26d9317 (cambios)
)

func main() {
	router := gin.Default()
	router.GET("/formulario", controllers.GetFormulario)
	router.POST("/formulario", controllers.PostFormulario)
	router.GET("/formulario/:id", controllers.GetFormularioByID)
	router.Run("localhost:8080")
<<<<<<< HEAD
}
=======
}
>>>>>>> 26d9317 (cambios)
