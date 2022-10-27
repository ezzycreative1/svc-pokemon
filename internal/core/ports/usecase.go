package ports

import (
	"context"

	"github.com/ezzycreative1/svc-pokemon/internal/core/domain"
)

type IRolesUsecase interface {
	FetchRoles(ctx context.Context) ([]domain.Roles, error)
	GetRoleByID(ctx context.Context, id int64) (*domain.Roles, error)
	UpdateRole(ctx context.Context, input *domain.Roles) error
	StoreRole(context.Context, *domain.Roles) error
	DeleteRole(ctx context.Context, id int64) error
}

type IUsersUsecase interface {
	FetchUsers(ctx context.Context) ([]domain.Users, error)
	GetUserByID(ctx context.Context, id int64) (*domain.Users, error)
	UpdateUser(ctx context.Context, input *domain.Users) error
	StoreUser(context.Context, *domain.Users) error
	DeleteUser(ctx context.Context, id int64) error
}

type IPokemonsUsecase interface {
	FetchPokemons(ctx context.Context) ([]domain.Pokemons, error)
	GetPokemonByID(ctx context.Context, id int64) (*domain.Pokemons, error)
	UpdatePokemon(ctx context.Context, input *domain.Pokemons) error
	StorePokemon(context.Context, *domain.Pokemons) error
	DeletePokemon(ctx context.Context, id int64) error
}

type IBattlesUsecase interface {
	FetchBatlles(ctx context.Context) ([]domain.Battles, error)
	GetBattleByID(ctx context.Context, id int64) (*domain.Battles, error)
	UpdateBattle(ctx context.Context, input *domain.Battles) error
	StoreBattle(context.Context, *domain.Battles) error
	DeleteBattle(ctx context.Context, id int64) error
}