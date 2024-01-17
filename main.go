package main


import
	"github.com/gin-gonic/gin"
	"github.com/SergioDimas17/controllers"



func main() {
	router := gin.Default()
	router.GET("/formulario", controllers.GetFormulario)
	router.POST("/formulario", controllers.PostFormulario)
	router.GET("/formulario/:id", controllers.GetFormularioByID)
	router.Run("localhost:8080")

}

