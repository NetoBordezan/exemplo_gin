package handlers

import (
	"api-gin/models"
	"api-gin/services"
	"api-gin/store"
	"github.com/gin-gonic/gin"
)

func AlocarSala(c *gin.Context) {
	tid := c.Param("id")
	t, ok := store.Turmas[tid]
	if !ok {
		c.JSON(404, gin.H{"erro": "turma não encontrada"})
		return
	}
	var req models.AlocacaoRequest
	if c.ShouldBindJSON(&req) != nil {
		c.JSON(400, gin.H{"erro": "dados inválidos"})
		return
	}
	s, ok := store.Salas[req.SalaID]
	if !ok {
		c.JSON(404, gin.H{"erro": "sala não encontrada"})
		return
	}
	if !s.Ativa || !t.Ativa {
		c.JSON(422, gin.H{"erro": "sala ou turma inativa"})
		return
	}
	if len(t.Alunos) > s.Capacidade {
		c.JSON(422, gin.H{"erro": "capacidade insuficiente"})
		return
	}
	if !services.DiaSemanaValido(req.DiaSemana) || !services.HorarioValido(req.HorarioInicio) || !services.HorarioValido(req.HorarioFim) || req.HorarioInicio >= req.HorarioFim {
		c.JSON(400, gin.H{"erro": "dia e horários HH:MM válidos são obrigatórios"})
		return
	}
	for _, o := range store.Turmas {
		if o.ID == t.ID || o.Alocacao == nil {
			continue
		}
		if o.Alocacao.SalaID == req.SalaID && o.Alocacao.DiaSemana == req.DiaSemana && services.HorariosSobrepostos(req.HorarioInicio, req.HorarioFim, o.Alocacao.HorarioInicio, o.Alocacao.HorarioFim) {
			c.JSON(409, gin.H{"erro": "conflito de agenda da sala"})
			return
		}
		if o.Alocacao.DiaSemana == req.DiaSemana && services.HorariosSobrepostos(req.HorarioInicio, req.HorarioFim, o.Alocacao.HorarioInicio, o.Alocacao.HorarioFim) {
			for _, a := range t.Alunos {
				for _, b := range o.Alunos {
					if a == b {
						c.JSON(409, gin.H{"erro": "conflito de agenda do aluno"})
						return
					}
				}
			}
		}
	}
	t.Alocacao = &models.Alocacao{SalaID: req.SalaID, DiaSemana: req.DiaSemana, HorarioInicio: req.HorarioInicio, HorarioFim: req.HorarioFim}
	store.Turmas[tid] = t
	c.JSON(200, gin.H{"mensagem": "turma alocada com sucesso", "alocacao": t.Alocacao})
}
