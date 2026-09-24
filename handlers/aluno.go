package handlers

import (
	"api-gin/models"
	"api-gin/store"
	"github.com/gin-gonic/gin"
)

func CriarAluno(c *gin.Context) {
	var x models.Aluno
	if c.ShouldBindJSON(&x) != nil || x.ID == "" || x.Nome == "" || x.Email == "" {
		c.JSON(400, gin.H{"erro": "id, nome e email são obrigatórios"})
		return
	}
	if _, ok := store.Alunos[x.ID]; ok {
		c.JSON(409, gin.H{"erro": "aluno já cadastrado"})
		return
	}
	store.Alunos[x.ID] = x
	c.JSON(201, x)
}
func ListarAlunos(c *gin.Context) {
	r := []models.Aluno{}
	for _, x := range store.Alunos {
		r = append(r, x)
	}
	c.JSON(200, r)
}
func BuscarAluno(c *gin.Context) {
	x, ok := store.Alunos[c.Param("id")]
	if !ok {
		c.JSON(404, gin.H{"erro": "aluno não encontrado"})
		return
	}
	c.JSON(200, x)
}
