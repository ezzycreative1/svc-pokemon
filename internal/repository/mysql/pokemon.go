package mysql

import (
	"context"
	"errors"
	"time"

	domain "github.com/ezzycreative1/svc-pokemon/internal/core/domain"
	"github.com/ezzycreative1/svc-pokemon/internal/data/customtype"
	"gorm.io/gorm"
)

const (
	conTimeout = 60

	// KeyTransaction concrete type for key context value transaction
	KTrans customtype.KeyTrx = customtype.KeyTrx("pokemon-trx")
)

type mysqlPokemonRepo struct {
	DB *gorm.DB
}

func NewMysqlPokemonRepo(db *gorm.DB) mysqlPokemonRepo {
	return mysqlPokemonRepo{
		DB: db,
	}
}

func (pr *mysqlPokemonRepo) FetchPokemons(ctx context.Context) ([]domain.Pokemons, error) {
	ctxWT, cancel := context.WithTimeout(ctx, conTimeout*time.Second)
	defer cancel()

	var res []domain.Pokemons
	query := pr.DB.WithContext(ctxWT).Find(&res)
	if query.Error != nil {
		return nil, query.Error
	}

	return res, nil
}

func (pr *mysqlPokemonRepo) StorePokemon(ctx context.Context, input *domain.Pokemons) error {
	trx, ok := ctx.Value(KeyTransaction).(*gorm.DB)
	if !ok {
		trx = pr.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, connectTimeout*time.Second)
	defer cancel()

	query := trx.WithContext(ctxWT).Create(&input)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (pr *mysqlPokemonRepo) UpdatePokemon(ctx context.Context, input *domain.Pokemons) error {
	trx, ok := ctx.Value(KeyTransaction).(*gorm.DB)
	if !ok {
		trx = pr.DB
	}

	if input.ID == 0 {
		return errors.New("pokemon to update must have id")
	}

	ctxWT, cancel := context.WithTimeout(ctx, connectTimeout*time.Second)
	defer cancel()

	query := trx.WithContext(ctxWT).Save(&input)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (pr *mysqlPokemonRepo) CheckExistsPokemon(ctx context.Context, name string) bool {
	ctxWT, cancel := context.WithTimeout(ctx, connectTimeout*time.Second)
	defer cancel()

	var exists bool
	if query := pr.DB.WithContext(ctxWT).Select("count(id) > 0").
		Where("name = ?", name).
		Find(&exists); query.Error != nil {
		return false
	}

	return true
}
