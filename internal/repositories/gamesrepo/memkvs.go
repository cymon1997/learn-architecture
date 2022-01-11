package gamesrepo

import (
	"encoding/json"
	"errors"

	"github.com/cymon1997/go-architecture/internal/core/ports"

	"github.com/cymon1997/go-architecture/internal/core/domain"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() ports.GamesRepository {
	return &memkvs{kvs: map[string][]byte{}}
}

func (repo *memkvs) Get(id string) (domain.Game, error) {
	if value, ok := repo.kvs[id]; ok {
		game := domain.Game{}
		err := json.Unmarshal(value, &game)
		if err != nil {
			return domain.Game{}, errors.New("fail to get value from kvs")
		}
		return game, nil
	}
	return domain.Game{}, errors.New("game not found in kvs")
}

func (repo *memkvs) Save(domain.Game) error {
	// TODO: Implement save
	return nil
}
