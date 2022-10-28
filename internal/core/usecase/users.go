package usecase

import (
	"context"
	"time"

	"github.com/ezzycreative1/svc-pokemon/internal/core/domain"
	"github.com/ezzycreative1/svc-pokemon/internal/core/ports"
	"github.com/ezzycreative1/svc-pokemon/pkg/errs"
	"github.com/ezzycreative1/svc-pokemon/pkg/mid"
	"github.com/ezzycreative1/svc-pokemon/pkg/mvalidator"
)

type userUseCase struct {
	Repo ports.IUsersRepository
}

func NewUserUsecase(repo ports.IUsersRepository) userUseCase {
	return userUseCase{
		Repo: repo,
	}
}

func (ru *userUseCase) LoginUser(ctx context.Context, input *domain.LoginRequest) (*domain.LoginResponse, error) {
	if input.Email == "" && input.Password == "" {
		return nil, errs.ErrBadParamInput
	}

	checkEmail := mvalidator.ValidEmail(input.Email)
	if !checkEmail {
		return nil, errs.ErrEmailWrong
	}

	user, _ := ru.Repo.GetUserByEmail(ctx, input.Email)
	checkPassword := mid.CheckPasswordHash(input.Password, user.Password)
	if !checkPassword {
		// check key exis
		return nil, errs.ErrBadParamInput
	}

	getToken, err := mid.GenerateToken(user.Email, user.ID)
	if err != nil {
		return nil, errs.ErrGenerateToken
	}
	//res.Token = getToken

	res := &domain.LoginResponse{
		Token: getToken,
	}

	return res, nil
}

func (ru *userUseCase) FetchUsers(ctx context.Context) ([]domain.UserResponse, error) {
	listUser, err := ru.Repo.FetchUsers(ctx)
	if err != nil {
		return nil, err
	}

	var res []domain.UserResponse
	for _, user := range listUser {
		res = append(res, domain.UserResponse{
			ID:     user.ID,
			Name:   user.Name,
			Email:  user.Email,
			Status: user.Status,
		})
	}
	return res, nil
}

func (ru *userUseCase) GetUserByID(ctx context.Context, id int64) (*domain.SingleUserResponse, error) {
	dataUser, err := ru.Repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	res := domain.SingleUserResponse{
		ID:        dataUser.ID,
		Role:      dataUser.RoleID,
		Name:      dataUser.Name,
		Email:     dataUser.Email,
		Status:    dataUser.Status,
		CreatedAt: dataUser.CreatedAt,
		UpdatedAt: dataUser.UpdatedAt,
	}
	return &res, nil
}

func (ru *userUseCase) UpdateUser(ctx context.Context, id int64, input *domain.UpdateUserRequest) error {
	dataUser, err := ru.Repo.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	data := domain.Users{
		ID:        dataUser.ID,
		RoleID:    input.Role,
		Name:      input.Fullname,
		Email:     input.Email,
		Password:  dataUser.Password,
		Status:    dataUser.Status,
		CreatedAt: dataUser.CreatedAt,
		UpdatedAt: time.Now(),
	}

	return ru.Repo.UpdateUser(ctx, &data)
}

func (ru *userUseCase) StoreUser(ctx context.Context, input *domain.StoreUserRequest) error {
	password, _ := mid.HashPassword(input.Password)

	data := domain.Users{
		RoleID:    input.Role,
		Name:      input.Fullname,
		Email:     input.Email,
		Password:  password,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := ru.Repo.StoreUser(ctx, &data); err != nil {
		return err
	}

	return nil
}

func (ru *userUseCase) DeleteUser(ctx context.Context, id int64) error {
	return ru.Repo.DeleteUser(ctx, id)
}
