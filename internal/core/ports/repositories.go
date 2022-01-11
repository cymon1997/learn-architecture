package ports

import "github.com/cymon1997/go-architecture/internal/core/domain"

type GamesRepository interface {
	Get(id string) (domain.Game, error)
	Save(domain.Game) error
}