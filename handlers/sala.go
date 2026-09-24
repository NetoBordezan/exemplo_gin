package handlers

import (
	"api-gin/models"
	"api-gin/store"

	"github.com/gin-gonic/gin"
)

func CriarSala(c *gin.Context) {
	var x models.Sala
	if c.ShouldBindJSON(&x) != nil || x.ID == "" || x.Nome == "" || x.Capacidade <= 0 {
		c.JSON(400, gin.H{"erro": "id, nome e capacidade maior que zero são obrigatórios"})
		return
	}
	if _, ok := store.Salas[x.ID]; ok {
		c.JSON(409, gin.H{"erro": "sala já cadastrada"})
		return
	}
	x.Ativa = true
	store.Salas[x.ID] = x
	c.JSON(201, x)
}
func ListarSalas(c *gin.Context) {
	r := []models.Sala{}
	for _, x := range store.Salas {
		r = append(r, x)
	}
	c.JSON(200, r)
}
func GradeSala(c *gin.Context) {
	id := c.Param("id")
	if _, ok := store.Salas[id]; !ok {
		c.JSON(404, gin.H{"erro": "sala não encontrada"})
		return
	}
	r := []gin.H{}
	for _, t := range store.Turmas {
		if t.Alocacao != nil && t.Alocacao.SalaID == id {
			r = append(r, gin.H{"turma_id": t.ID, "turma": t.Nome, "dia_semana": t.Alocacao.DiaSemana, "horario_inicio": t.Alocacao.HorarioInicio, "horario_fim": t.Alocacao.HorarioFim})
		}
	}
	c.JSON(200, r)
}
