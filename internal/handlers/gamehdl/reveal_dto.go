package gamehdl

import "github.com/cymon1997/learn-architecture/internal/core/domain"

type BodyRevealCell struct {
	Row uint `json:"row"`
	Col uint `json:"col"`
}

type ResponseRevealCell domain.Game

func BuildResponseRevealCell(model domain.Game) ResponseRevealCell {
	return ResponseRevealCell(model)
}
