package external

import (
	"context"

	"github.com/ezzycreative1/svc-pokemon/internal/core/domain"
)

type externalRepo struct {
}

func NewExternalRepo() externalRepo {
	return externalRepo{}
}

func (ex *externalRepo) GetPokemon(ctx context.Context) ([]domain.PokemonExternal, error) {
	return nil, nil
}
