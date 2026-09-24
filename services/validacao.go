package services

import (
	"strings"
	"time"
)

func HorarioValido(h string) bool {
	_, err := time.Parse("15:04", h)
	return err == nil
}

func DiaSemanaValido(dia string) bool {
	dia = strings.ToLower(strings.TrimSpace(dia))
	switch dia {
	case "segunda", "segunda-feira", "terça", "terca", "terça-feira", "terca-feira",
		"quarta", "quarta-feira", "quinta", "quinta-feira", "sexta", "sexta-feira",
		"sábado", "sabado", "domingo":
		return true
	default:
		return false
	}
}

func HorariosSobrepostos(inicioNovo, fimNovo, inicioExistente, fimExistente string) bool {
	return inicioNovo < fimExistente && fimNovo > inicioExistente
}
