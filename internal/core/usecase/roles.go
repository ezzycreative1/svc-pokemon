package usecase

import (
	"context"
	"time"

	"github.com/ezzycreative1/svc-pokemon/internal/core/domain"
	"github.com/ezzycreative1/svc-pokemon/internal/core/ports"
)

type rolesUseCase struct {
	Repo ports.IRolesRepository
}

func NewRolesUsecase(repo ports.IRolesRepository) rolesUseCase {
	return rolesUseCase{
		Repo: repo,
	}
}

func (ru *rolesUseCase) FetchRoles(ctx context.Context) ([]domain.Roles, error) {
	res, err := ru.Repo.FetchRoles(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ru *rolesUseCase) GetRoleByID(ctx context.Context, id int64) (*domain.Roles, error) {
	res, err := ru.Repo.GetRoleByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ru *rolesUseCase) UpdateRole(ctx context.Context, input *domain.Roles) error {
	input.UpdatedAt = time.Now()
	return ru.Repo.UpdateRole(ctx, input)
}

func (ru *rolesUseCase) StoreRole(ctx context.Context, input *domain.Roles) error {
	err := ru.Repo.StoreRole(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (ru *rolesUseCase) DeleteRole(ctx context.Context, id int64) error {
	return ru.Repo.DeleteRole(ctx, id)
}
