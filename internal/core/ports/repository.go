package ports

import (
	"context"

	"github.com/ezzycreative1/svc-pokemon/internal/core/domain"
)

type IRolesRepository interface {
	FetchRoles(ctx context.Context) ([]domain.Roles, error)
	GetRoleByID(ctx context.Context, id int64) (*domain.Roles, error)
	UpdateRole(ctx context.Context, input *domain.Roles) error
	StoreRole(ctx context.Context, input *domain.Roles) error
	DeleteRole(ctx context.Context, id int64) error
	GetRoleID(ctx context.Context, name string) (int64, error)
}

type IUsersRepository interface {
	FetchUsers(ctx context.Context) ([]domain.Users, error)
	GetUserByID(ctx context.Context, id int64) (*domain.Users, error)
	UpdateUser(ctx context.Context, input *domain.Users) error
	StoreUser(ctx context.Context, input *domain.Users) error
	DeleteUser(ctx context.Context, id int64) error
	GetUserByEmail(ctx context.Context, email string) (*domain.Users, error)
}

type IPokemonsRepository interface {
	FetchPokemons(ctx context.Context) ([]domain.Pokemons, error)
	GetPokemonByID(ctx context.Context, id int64) (*domain.Pokemons, error)
	UpdatePokemon(ctx context.Context, input *domain.Pokemons) error
	StorePokemon(ctx context.Context, input *domain.Pokemons) error
	DeletePokemon(ctx context.Context, id int64) error
}

type IBattleRepository interface {
	FetchBattles(ctx context.Context) ([]domain.Pokemons, error)
	GetPokemonByID(ctx context.Context, id int64) (*domain.Pokemons, error)
	UpdatePokemon(ctx context.Context, input *domain.Pokemons) error
	StorePokemon(ctx context.Context, input *domain.Pokemons) error
	DeletePokemon(ctx context.Context, id int64) error
}
