<<<<<<< HEAD
package controllers
=======
// controllers/controllers.go

package main
>>>>>>> c354c43 (go)

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Formulario struct {
	ID      string `json:"id"`
	Nombre  string `json:"nombre"`
	Correo  string `json:"correo"`
	Telefono int    `json:"telefono"`
}

var Formularios = []Formulario{
	{ID: "1", Nombre: "Juan Cardenas", Correo: "juancardenas@gmail.com", Telefono: 1234567},
	{ID: "2", Nombre: "Sara Giraldo", Correo: "SaraGiraldo@gmail.com", Telefono: 3456345},
	{ID: "3", Nombre: "Diego Rovira", Correo: "DiegoRovira@gmail.com", Telefono: 4357654},
	{ID: "4", Nombre: "Sofia Argel", Correo: "SofiaArgel@gmail.com", Telefono: 765564},
	{ID: "5", Nombre: "Luis Quintero", Correo: "LuisQuintero@gmail.com", Telefono: 23214567},
}

func GetFormulario(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Formularios)
}

func PostFormulario(c *gin.Context) {
	var newFormulario Formulario

	if err := c.BindJSON(&newFormulario); err != nil {
		return
	}

	Formularios = append(Formularios, newFormulario)
	c.IndentedJSON(http.StatusCreated, newFormulario)
}

func GetFormularioByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range Formularios {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Datos no encontrado"})
}
