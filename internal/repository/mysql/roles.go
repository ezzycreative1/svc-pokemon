package mysql

import (
	"context"
	"errors"
	"time"

	domain "github.com/ezzycreative1/svc-pokemon/internal/core/domain"
	"github.com/ezzycreative1/svc-pokemon/internal/data/customtype"
	"github.com/ezzycreative1/svc-pokemon/pkg/errs"
	"gorm.io/gorm"
)

const (
	connectTimeout = 60

	// KeyTransaction concrete type for key context value transaction
	KeyTransaction customtype.KeyTrx = customtype.KeyTrx("pokemon-trx")
)

type mysqlRolesRepo struct {
	DB *gorm.DB
}

func NewMysqlRolesRepo(db *gorm.DB) mysqlRolesRepo {
	return mysqlRolesRepo{
		DB: db,
	}
}

func (cr *mysqlRolesRepo) FetchRoles(ctx context.Context) ([]domain.Roles, error) {
	ctxWT, cancel := context.WithTimeout(ctx, connectTimeout*time.Second)
	defer cancel()

	var res []domain.Roles
	query := cr.DB.WithContext(ctxWT).Find(&res)
	if query.Error != nil {
		return nil, query.Error
	}

	return res, nil
}

func (cr *mysqlRolesRepo) GetRoleByID(ctx context.Context, id int64) (*domain.Roles, error) {
	trx, ok := ctx.Value(KeyTransaction).(*gorm.DB)
	if !ok {
		trx = cr.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, connectTimeout*time.Second)
	defer cancel()

	var result domain.Roles
	query := trx.WithContext(ctxWT).
		Where("id = ?", id).
		Limit(1).
		Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	if result.ID == 0 {
		return nil, errs.ErrNotFound
	}

	return &result, nil

}

func (cr *mysqlRolesRepo) StoreRole(ctx context.Context, input *domain.Roles) error {
	trx, ok := ctx.Value(KeyTransaction).(*gorm.DB)
	if !ok {
		trx = cr.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, connectTimeout*time.Second)
	defer cancel()

	query := trx.WithContext(ctxWT).Create(&input)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (cr *mysqlRolesRepo) DeleteRole(ctx context.Context, id int64) error {
	trx, ok := ctx.Value(KeyTransaction).(*gorm.DB)
	if !ok {
		trx = cr.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, connectTimeout*time.Second)
	defer cancel()

	query := trx.WithContext(ctxWT).Delete(&domain.Roles{}, id)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (cr *mysqlRolesRepo) UpdateRole(ctx context.Context, input *domain.Roles) error {
	trx, ok := ctx.Value(KeyTransaction).(*gorm.DB)
	if !ok {
		trx = cr.DB
	}

	if input.ID == 0 {
		return errors.New("roles to update must have id")
	}

	ctxWT, cancel := context.WithTimeout(ctx, connectTimeout*time.Second)
	defer cancel()

	query := trx.WithContext(ctxWT).Save(&input)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (cr *mysqlRolesRepo) GetRoleID(ctx context.Context, name string) (int64, error) {
	trx, ok := ctx.Value(KeyTransaction).(*gorm.DB)
	if !ok {
		trx = cr.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, connectTimeout*time.Second)
	defer cancel()

	var result domain.Roles
	query := trx.WithContext(ctxWT).
		Select("id").
		Where("name = ?", name).
		Limit(1).
		Find(&result)

	if query.Error != nil {
		return 0, query.Error
	}

	return result.ID, nil

}
