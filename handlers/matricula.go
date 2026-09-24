package handlers

import (
	"api-gin/models"
	"api-gin/services"
	"api-gin/store"
	"github.com/gin-gonic/gin"
)

func MatricularAluno(c *gin.Context) {
	tid := c.Param("id")
	t, ok := store.Turmas[tid]
	if !ok {
		c.JSON(404, gin.H{"erro": "turma não encontrada"})
		return
	}
	var req models.MatriculaRequest
	if c.ShouldBindJSON(&req) != nil || req.AlunoID == "" {
		c.JSON(400, gin.H{"erro": "aluno_id é obrigatório"})
		return
	}
	if _, ok := store.Alunos[req.AlunoID]; !ok {
		c.JSON(404, gin.H{"erro": "aluno não encontrado"})
		return
	}
	for _, id := range t.Alunos {
		if id == req.AlunoID {
			c.JSON(409, gin.H{"erro": "aluno já matriculado na turma"})
			return
		}
	}
	if t.Alocacao != nil {
		if len(t.Alunos)+1 > store.Salas[t.Alocacao.SalaID].Capacidade {
			c.JSON(422, gin.H{"erro": "capacidade insuficiente"})
			return
		}
		for _, o := range store.Turmas {
			if o.ID == t.ID || o.Alocacao == nil || o.Alocacao.DiaSemana != t.Alocacao.DiaSemana {
				continue
			}
			mat := false
			for _, id := range o.Alunos {
				if id == req.AlunoID {
					mat = true
					break
				}
			}
			if mat && services.HorariosSobrepostos(t.Alocacao.HorarioInicio, t.Alocacao.HorarioFim, o.Alocacao.HorarioInicio, o.Alocacao.HorarioFim) {
				c.JSON(409, gin.H{"erro": "conflito de agenda do aluno"})
				return
			}
		}
	}
	t.Alunos = append(t.Alunos, req.AlunoID)
	store.Turmas[tid] = t
	c.JSON(200, gin.H{"mensagem": "aluno matriculado com sucesso"})
}
func ListarAlunosTurma(c *gin.Context) {
	t, ok := store.Turmas[c.Param("id")]
	if !ok {
		c.JSON(404, gin.H{"erro": "turma não encontrada"})
		return
	}
	r := []models.Aluno{}
	for _, id := range t.Alunos {
		r = append(r, store.Alunos[id])
	}
	c.JSON(200, r)
}
