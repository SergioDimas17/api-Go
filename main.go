package main

<<<<<<< HEAD

import
	"github.com/gin-gonic/gin"
	"github.com/SergioDimas17/controllers"


=======
import (
	"github.com/SergioDimas17/webservicegin/controllers"
    "github.com/gin-gonic/gin"
)
>>>>>>> c354c43 (go)

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
>>>>>>> c354c43 (go)
