package handlers

import (
	"api-gin/models"
	"api-gin/store"
	"github.com/gin-gonic/gin"
)

func CriarTurma(c *gin.Context) {
	var x models.Turma
	if c.ShouldBindJSON(&x) != nil || x.ID == "" || x.Nome == "" || x.Disciplina == "" || x.DocenteID == "" {
		c.JSON(400, gin.H{"erro": "id, nome, disciplina e docente_id são obrigatórios"})
		return
	}
	if _, ok := store.Turmas[x.ID]; ok {
		c.JSON(409, gin.H{"erro": "turma já cadastrada"})
		return
	}
	if _, ok := store.Docentes[x.DocenteID]; !ok {
		c.JSON(404, gin.H{"erro": "docente não encontrado"})
		return
	}
	x.Ativa = true
	x.Alunos = []string{}
	store.Turmas[x.ID] = x
	c.JSON(201, x)
}
func ListarTurmas(c *gin.Context) {
	r := []gin.H{}
	for _, t := range store.Turmas {
		r = append(r, gin.H{"id": t.ID, "nome": t.Nome, "disciplina": t.Disciplina, "docente_id": t.DocenteID, "quantidade_alunos": len(t.Alunos), "alocada": t.Alocacao != nil, "alocacao": t.Alocacao})
	}
	c.JSON(200, r)
}
