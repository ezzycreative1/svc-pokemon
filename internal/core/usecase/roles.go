package usecase

import (
	"context"
	"time"

	"github.com/ezzycreative1/svc-pokemon/internal/core/domain"
	"github.com/ezzycreative1/svc-pokemon/internal/core/ports"
)

const seq = 1

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

func (ru *rolesUseCase) UpdateRole(ctx context.Context, id int64, input *domain.RoleRequest) error {
	dataRole, err := ru.Repo.GetRoleByID(ctx, id)
	if err != nil {
		return err
	}

	data := domain.Roles{
		ID:        dataRole.ID,
		Name:      input.Name,
		Status:    dataRole.Status,
		CreatedAt: dataRole.CreatedAt,
		UpdatedAt: time.Now(),
	}

	return ru.Repo.UpdateRole(ctx, &data)
}

func (ru *rolesUseCase) StoreRole(ctx context.Context, input *domain.RoleRequest) error {
	data := domain.Roles{
		Name:      input.Name,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := ru.Repo.StoreRole(ctx, &data)
	if err != nil {
		return err
	}

	return nil
}

func (ru *rolesUseCase) DeleteRole(ctx context.Context, id int64) error {
	data, err := ru.Repo.GetRoleByID(ctx, id)
	if err != nil {
		return err
	}
	return ru.Repo.DeleteRole(ctx, data.ID)
}

func getRoleName(id int64) string {
	var roleName string
	switch id {
	case 1:
		roleName = "Bos"
	case 2:
		roleName = "Admin"
	case 3:
		roleName = "Operasional"
	case 4:
		roleName = "BlackDealer"
	}
	return roleName
}

func (ru *rolesUseCase) GetRoleID(ctx context.Context, name string) (int64, error) {
	return ru.Repo.GetRoleID(ctx, name)
}
