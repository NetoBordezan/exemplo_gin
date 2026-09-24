package main

import (
	"api-gin/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", handlers.Health)
		v1.POST("/salas", handlers.CriarSala)
		v1.GET("/salas", handlers.ListarSalas)
		v1.GET("/salas/:id/grade", handlers.GradeSala)
		v1.POST("/alunos", handlers.CriarAluno)
		v1.GET("/alunos", handlers.ListarAlunos)
		v1.GET("/alunos/:id", handlers.BuscarAluno)
		v1.POST("/docentes", handlers.CriarDocente)
		v1.GET("/docentes", handlers.ListarDocentes)
		v1.GET("/docentes/:id", handlers.BuscarDocente)
		v1.POST("/turmas", handlers.CriarTurma)
		v1.GET("/turmas", handlers.ListarTurmas)
		v1.POST("/turmas/:id/alunos", handlers.MatricularAluno)
		v1.GET("/turmas/:id/alunos", handlers.ListarAlunosTurma)
		v1.POST("/turmas/:id/alocar", handlers.AlocarSala)
	}
	r.Run(":8080")
}
