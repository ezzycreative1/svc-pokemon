package usecase

import (
	"context"
	"time"

	"github.com/ezzycreative1/svc-pokemon/internal/core/domain"
	"github.com/ezzycreative1/svc-pokemon/internal/core/ports"
)

type pokemonUseCase struct {
	Repo   ports.IPokemonsRepository
	ExRepo ports.IPokemonExternalRepository
}

func NewPokemonUsecase(repo ports.IPokemonsRepository, exRepo ports.IPokemonExternalRepository) pokemonUseCase {
	return pokemonUseCase{
		Repo:   repo,
		ExRepo: exRepo,
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
			ID:   poke.ID,
			Name: poke.Name,
			//UserID:    poke.UserID,
			Stock:     poke.Stock,
			CreatedAt: poke.CreatedAt,
			UpdatedAt: poke.UpdatedAt,
		})
	}
	return res, nil
}

func (p *pokemonUseCase) StorePokemon(ctx context.Context, input *domain.StorePokemonRequest) error {
	dataPokemon, err := p.ExRepo.GetPokemonByName(ctx, input.Name)
	if err != nil {
		return err
	}

	data := &domain.Pokemons{
		Name: dataPokemon.Name,
		//UserID:    input.UserID,
		Stock:     input.Stock,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// if (p.Repo.CheckExistsPokemon(ctx, input.Name) {

	// }

	if err := p.Repo.StorePokemon(ctx, data); err != nil {
		return err
	}
	return nil
}
