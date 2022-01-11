package gamesrv

import (
	"errors"

	"github.com/cymon1997/learn-architecture/internal/core/domain"
	"github.com/cymon1997/learn-architecture/internal/core/ports"
	"github.com/google/uuid"
)

type service struct {
	gamesRepository ports.GamesRepository
}

func New(gamesRepository ports.GamesRepository) *service {
	return &service{
		gamesRepository: gamesRepository,
	}
}

func (srv *service) Get(id string) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		return domain.Game{}, errors.New(err.Error())
	}
	game.Board = game.Board.HideBombs()
	return game, nil
}

func (srv *service) Create(name string, size uint, bombs uint) (domain.Game, error) {
	if bombs >= size*size {
		return domain.Game{}, errors.New("the number of bombs is too high")
	}
	game := domain.NewGame(uuid.New().String(), name, size, bombs)
	if err := srv.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New("create game into repository has failed")
	}
	game.Board = game.Board.HideBombs()
	return game, nil
}

func (srv *service) Reveal(id string, row uint, col uint) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		return domain.Game{}, errors.New(err.Error())
	}
	if !game.Board.IsValidPosition(row, col) {
		return domain.Game{}, errors.New("invalid position")
	}
	if game.IsOver() {
		return domain.Game{}, errors.New("game is over")
	}
	if game.Board.Contains(row, col, domain.CELL_BOMB) {
		game.State = domain.GAME_STATE_LOST
	} else {
		game.Board.Set(row, col, domain.CELL_REVEALED)
		if !game.Board.HasEmptyCells() {
			game.State = domain.GAME_STATE_WON
		}
	}
	if err := srv.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New("update game into repository has failed")
	}
	game.Board = game.Board.HideBombs()
	return game, nil
}
