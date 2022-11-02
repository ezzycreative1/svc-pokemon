package usecase

import (
	"context"
	"time"

	"github.com/ezzycreative1/svc-pokemon/internal/core/domain"
	"github.com/ezzycreative1/svc-pokemon/internal/core/ports"
)

type pokemonTypeUseCase struct {
	Repo       ports.IPokemonTypeRepository
	ExTypeRepo ports.IPokemonTypeExternalRepository
}

func NewPokemonTypeUsecase(repo ports.IPokemonTypeRepository, exRepo ports.IPokemonTypeExternalRepository) pokemonTypeUseCase {
	return pokemonTypeUseCase{
		Repo:       repo,
		ExTypeRepo: exRepo,
	}
}

func (p *pokemonTypeUseCase) FetchPokemonTypes(ctx context.Context) ([]domain.PokemonTypes, error) {
	listPokemon, err := p.Repo.FetchPokemonTypes(ctx)
	if err != nil {
		return nil, err
	}

	var res []domain.PokemonTypes
	for _, poke := range listPokemon {
		res = append(res, domain.PokemonTypes{
			ID:        poke.ID,
			Name:      poke.Name,
			CreatedAt: poke.CreatedAt,
			UpdatedAt: poke.UpdatedAt,
		})
	}
	return res, nil
}

func (p *pokemonUseCase) StorePokemonType(ctx context.Context, input *domain.StorePokemonRequest) error {
	dataPokemon, err := p.ExRepo.GetPokemonByName(ctx, input.Name)
	if err != nil {
		return err
	}
	data := &domain.Pokemons{
		Name:      dataPokemon.Name,
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
