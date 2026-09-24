package models

type Sala struct {
	ID         string   `json:"id"`
	Nome       string   `json:"nome"`
	Capacidade int      `json:"capacidade"`
	Recursos   []string `json:"recursos"`
	Ativa      bool     `json:"ativa"`
}
type Aluno struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
type Docente struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
type Alocacao struct {
	SalaID        string `json:"sala_id"`
	DiaSemana     string `json:"dia_semana"`
	HorarioInicio string `json:"horario_inicio"`
	HorarioFim    string `json:"horario_fim"`
}
type Turma struct {
	ID         string    `json:"id"`
	Nome       string    `json:"nome"`
	Disciplina string    `json:"disciplina"`
	DocenteID  string    `json:"docente_id"`
	Alunos     []string  `json:"alunos"`
	Ativa      bool      `json:"ativa"`
	Alocacao   *Alocacao `json:"alocacao,omitempty"`
}
type MatriculaRequest struct {
	AlunoID string `json:"aluno_id"`
}
type AlocacaoRequest struct {
	SalaID        string `json:"sala_id"`
	DiaSemana     string `json:"dia_semana"`
	HorarioInicio string `json:"horario_inicio"`
	HorarioFim    string `json:"horario_fim"`
}
