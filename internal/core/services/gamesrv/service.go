package gamesrv

import (
	"errors"

	"github.com/cymon1997/go-architecture/internal/core/domain"
	"github.com/cymon1997/go-architecture/internal/core/ports"
)

type service struct {
	gamesRepository ports.GamesRepository
}

func New(gamesRepository ports.GamesRepository) ports.GamesService {
	return &service{
		gamesRepository: gamesRepository,
	}
}

func (srv *service) Get(id string) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		return domain.Game{}, errors.New("get game from repository has failed")
	}
	return game, nil
}

func (srv *service) Create(name string, size uint, bombs uint) (domain.Game, error) {
	if bombs >= size*size {
		return domain.Game{}, errors.New("the number of bombs is invalid")
	}
	game := domain.NewGame(name, size, bombs)
	if err := srv.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New("create game into repository has failed")
	}
	return game, nil
}

func (srv *service) Reveal(id string, row uint, col uint) (domain.Game, error) {
	// TODO: Implement reveal
	return domain.Game{}, nil
}
