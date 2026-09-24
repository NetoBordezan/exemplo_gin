package store

import "api-gin/models"

var Salas = make(map[string]models.Sala)
var Alunos = make(map[string]models.Aluno)
var Docentes = make(map[string]models.Docente)
var Turmas = make(map[string]models.Turma)
