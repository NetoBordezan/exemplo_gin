package handlers

import (
	"api-gin/models"
	"api-gin/store"
	"github.com/gin-gonic/gin"
)

func CriarDocente(c *gin.Context) {
	var x models.Docente
	if c.ShouldBindJSON(&x) != nil || x.ID == "" || x.Nome == "" || x.Email == "" {
		c.JSON(400, gin.H{"erro": "id, nome e email são obrigatórios"})
		return
	}
	if _, ok := store.Docentes[x.ID]; ok {
		c.JSON(409, gin.H{"erro": "docente já cadastrado"})
		return
	}
	store.Docentes[x.ID] = x
	c.JSON(201, x)
}
func ListarDocentes(c *gin.Context) {
	r := []models.Docente{}
	for _, x := range store.Docentes {
		r = append(r, x)
	}
	c.JSON(200, r)
}
func BuscarDocente(c *gin.Context) {
	x, ok := store.Docentes[c.Param("id")]
	if !ok {
		c.JSON(404, gin.H{"erro": "docente não encontrado"})
		return
	}
	c.JSON(200, x)
}
