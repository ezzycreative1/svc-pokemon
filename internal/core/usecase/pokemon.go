package usecase

import (
	"context"

	"github.com/ezzycreative1/svc-pokemon/internal/core/domain"
	"github.com/ezzycreative1/svc-pokemon/internal/core/ports"
)

type pokemonUseCase struct {
	Repo ports.IPokemonsRepository
}

func NewPokemonUsecase(repo ports.IPokemonsRepository) pokemonUseCase {
	return pokemonUseCase{
		Repo: repo,
	}
}

func (p *pokemonUseCase) FetchPokemons(ctx context.Context) ([]domain.Pokemons, error) {
	listPokemon, err := p.Repo.FetchPokemons(ctx)
	if err != nil {
		return nil, err
	}

	var res []domain.Pokemons
	for _, poke := range listPokemon {
		res = append(res, domain.Pokemons{
			ID:        poke.ID,
			Name:      poke.Name,
			UserID:    poke.UserID,
			Stock:     poke.Stock,
			CreatedAt: poke.CreatedAt,
			UpdatedAt: poke.UpdatedAt,
		})
	}
	return res, nil
}

func (p *pokemonUseCase) StorePokemon(ctx context.Context, input *domain.StorePokemonRequest) error {
	if err := p.Repo.StorePokemon(ctx, input); err != nil {
		return err
	}
	return nil
}
