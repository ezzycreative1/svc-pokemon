package usecase

import (
	"context"
	"time"

	"github.com/ezzycreative1/svc-pokemon/internal/core/domain"
	"github.com/ezzycreative1/svc-pokemon/internal/core/ports"
)

type userUseCase struct {
	Repo ports.IUsersRepository
}

func NewUserUsecase(repo ports.IUsersRepository) userUseCase {
	return userUseCase{
		Repo: repo,
	}
}

func (ru *userUseCase) FetchUsers(ctx context.Context) ([]domain.Users, error) {
	res, err := ru.Repo.FetchUsers(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ru *userUseCase) GetUserByID(ctx context.Context, id int64) (*domain.Users, error) {
	res, err := ru.Repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ru *userUseCase) UpdateUser(ctx context.Context, input *domain.Users) error {
	input.UpdatedAt = time.Now()
	return ru.Repo.UpdateUser(ctx, input)
}

func (ru *userUseCase) StoreUser(ctx context.Context, input *domain.Users) error {
	err := ru.Repo.StoreUser(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (ru *userUseCase) DeleteUser(ctx context.Context, id int64) error {
	return ru.Repo.DeleteUser(ctx, id)
}
