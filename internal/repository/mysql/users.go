package mysql

import (
	"context"
	"errors"
	"fmt"
	"time"

	domain "github.com/ezzycreative1/svc-pokemon/internal/core/domain"
	"github.com/ezzycreative1/svc-pokemon/internal/data/customtype"
	"github.com/ezzycreative1/svc-pokemon/pkg/errs"
	"gorm.io/gorm"
)

const (
	connTimeout = 60

	// KeyTransaction concrete type for key context value transaction
	KeyTrans customtype.KeyTrx = customtype.KeyTrx("pokemon-trx")
)

type mysqlUsersRepo struct {
	DB *gorm.DB
}

func NewMysqlUserRepo(db *gorm.DB) mysqlUsersRepo {
	return mysqlUsersRepo{
		DB: db,
	}
}

func (cr *mysqlUsersRepo) FetchUsers(ctx context.Context) ([]domain.Users, error) {
	ctxWT, cancel := context.WithTimeout(ctx, connTimeout*time.Second)
	defer cancel()

	var res []domain.Users
	query := cr.DB.WithContext(ctxWT).Find(&res)
	if query.Error != nil {
		return nil, query.Error
	}

	return res, nil
}

func (cr *mysqlUsersRepo) GetUserByID(ctx context.Context, id int64) (*domain.Users, error) {
	trx, ok := ctx.Value(KeyTrans).(*gorm.DB)
	if !ok {
		trx = cr.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, connTimeout*time.Second)
	defer cancel()

	var result domain.Users
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

func (cr *mysqlUsersRepo) StoreUser(ctx context.Context, input *domain.Users) error {
	fmt.Println(input)
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

func (cr *mysqlUsersRepo) DeleteUser(ctx context.Context, id int64) error {
	trx, ok := ctx.Value(KeyTransaction).(*gorm.DB)
	if !ok {
		trx = cr.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, connectTimeout*time.Second)
	defer cancel()

	query := trx.WithContext(ctxWT).Delete(&domain.Users{}, id)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (cr *mysqlUsersRepo) UpdateUser(ctx context.Context, input *domain.Users) error {
	trx, ok := ctx.Value(KeyTransaction).(*gorm.DB)
	if !ok {
		trx = cr.DB
	}

	if input.ID == 0 {
		return errors.New("users to update must have id")
	}

	ctxWT, cancel := context.WithTimeout(ctx, connectTimeout*time.Second)
	defer cancel()

	query := trx.WithContext(ctxWT).Save(&input)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (cr *mysqlUsersRepo) GetUserByEmail(ctx context.Context, email string) (*domain.Users, error) {
	trx, ok := ctx.Value(KeyTrans).(*gorm.DB)
	if !ok {
		trx = cr.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, connTimeout*time.Second)
	defer cancel()

	var result domain.Users
	query := trx.WithContext(ctxWT).
		Where("email = ?", email).
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
