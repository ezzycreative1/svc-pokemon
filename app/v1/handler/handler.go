package handler

import (
	"fmt"
	"net/http"
	"strconv"

	domain "github.com/ezzycreative1/svc-pokemon/internal/core/domain"

	"github.com/ezzycreative1/svc-pokemon/config"
	"github.com/ezzycreative1/svc-pokemon/internal/core/ports"
	"github.com/ezzycreative1/svc-pokemon/pkg/mid"
	"github.com/ezzycreative1/svc-pokemon/pkg/mlog"
	"github.com/ezzycreative1/svc-pokemon/pkg/mvalidator"
	"github.com/ezzycreative1/svc-pokemon/pkg/web"
	"github.com/labstack/echo/v4"
)

type PokemonHandler struct {
	UseCaseRoles ports.IRolesUsecase
	UseCaseUsers ports.IUsersUsecase
	Validator    mvalidator.Validator
	Logger       mlog.Logger
	Cfg          config.Group
}

func NewPokemonHandler(
	usecaseRoles ports.IRolesUsecase,
	usecaseUsers ports.IUsersUsecase,
	validator mvalidator.Validator,
	logger mlog.Logger,
	config config.Group,
) PokemonHandler {
	return PokemonHandler{
		UseCaseRoles: usecaseRoles,
		UseCaseUsers: usecaseUsers,
		Validator:    validator,
		Logger:       logger,
		Cfg:          config,
	}
}

//Handler Roles
func (ch *PokemonHandler) FetchRoles(ctx echo.Context) error {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Request().Context(), requestID)

	listAr, err := ch.UseCaseRoles.FetchRoles(userCtx)
	if err != nil {
		ch.Logger.ErrorT(requestID, "error fetch data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", listAr, nil)
}

func (ch *PokemonHandler) GetRoleByID(ctx echo.Context) error {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Request().Context(), requestID)

	idP, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ch.Logger.ErrorT(requestID, "error get data", err)
		return web.ResponseFormatter(ctx, http.StatusNotFound, err.Error(), nil, err)
	}

	id := int64(idP)

	art, err := ch.UseCaseRoles.GetRoleByID(userCtx, id)
	if err != nil {
		ch.Logger.ErrorT(requestID, "error get data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "success", art, nil)
}

func (ch *PokemonHandler) StoreRole(ctx echo.Context) (err error) {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Request().Context(), requestID)

	var payload domain.RoleRequest
	fmt.Println(payload)
	if err := ctx.Bind(&payload); err != nil {
		ch.Logger.ErrorT(requestID, "role store payload", err, mlog.Any("payload", payload))
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	mapErr, err := ch.Validator.Struct(payload)
	if err != nil {
		ch.Logger.ErrorT(requestID, "Bad Request", err)
		return web.ResponseErrValidation(ctx, "bad request", mapErr)
	}

	err = ch.UseCaseRoles.StoreRole(userCtx, &payload)
	if err != nil {
		ch.Logger.ErrorT(requestID, "error store data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", "", nil)
}

func (ch *PokemonHandler) UpdateRole(ctx echo.Context) (err error) {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Request().Context(), requestID)

	idP, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ch.Logger.ErrorT(requestID, "error get data", err)
		return web.ResponseFormatter(ctx, http.StatusNotFound, err.Error(), nil, err)
	}

	id := int64(idP)

	var payload domain.RoleRequest
	if err := ctx.Bind(&payload); err != nil {
		ch.Logger.ErrorT(requestID, "role store payload", err, mlog.Any("payload", payload))
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	mapErr, err := ch.Validator.Struct(payload)
	if err != nil {
		ch.Logger.ErrorT(requestID, "Bad Request", err)
		return web.ResponseErrValidation(ctx, "bad request", mapErr)
	}

	err = ch.UseCaseRoles.UpdateRole(userCtx, id, &payload)
	if err != nil {
		ch.Logger.ErrorT(requestID, "error update data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", "", nil)
}

func (ch *PokemonHandler) DeleteRole(c echo.Context) error {
	requestID := mid.GetID(c)
	userCtx := mid.SetIDx(c.Request().Context(), requestID)

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ch.Logger.ErrorT(requestID, "error delete data", err)
		return web.ResponseFormatter(c, http.StatusNotFound, err.Error(), nil, err)
	}

	id := int64(idP)

	err = ch.UseCaseRoles.DeleteRole(userCtx, id)
	if err != nil {
		ch.Logger.ErrorT(requestID, "error delete data", err)
		return web.ResponseFormatter(c, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(c, http.StatusNoContent, "success", "", nil)
}

//Handler User
func (ch *PokemonHandler) Login(ctx echo.Context) error {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Request().Context(), requestID)

	var payload domain.LoginRequest
	if err := ctx.Bind(&payload); err != nil {
		ch.Logger.ErrorT(requestID, "login user payload", err, mlog.Any("payload", payload))
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	mapErr, err := ch.Validator.Struct(payload)
	if err != nil {
		ch.Logger.ErrorT(requestID, "Bad Request", err)
		return web.ResponseErrValidation(ctx, "bad request", mapErr)
	}

	data, err := ch.UseCaseUsers.LoginUser(userCtx, &payload)
	if err != nil {
		ch.Logger.ErrorT(requestID, "error store data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", data, nil)
}

func (ch *PokemonHandler) FetchUsers(ctx echo.Context) error {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Request().Context(), requestID)

	listAr, err := ch.UseCaseUsers.FetchUsers(userCtx)
	if err != nil {
		ch.Logger.ErrorT(requestID, "error fetch data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", listAr, nil)
}

func (ch *PokemonHandler) GetUserByID(ctx echo.Context) error {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Request().Context(), requestID)

	idP, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ch.Logger.ErrorT(requestID, "error get data", err)
		return web.ResponseFormatter(ctx, http.StatusNotFound, err.Error(), nil, err)
	}

	id := int64(idP)

	art, err := ch.UseCaseUsers.GetUserByID(userCtx, id)
	if err != nil {
		ch.Logger.ErrorT(requestID, "error get data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "success", art, nil)
}

func (ch *PokemonHandler) StoreUser(ctx echo.Context) (err error) {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Request().Context(), requestID)

	var payload domain.StoreUserRequest
	if err := ctx.Bind(&payload); err != nil {
		ch.Logger.ErrorT(requestID, "user store payload", err, mlog.Any("payload", payload))
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	mapErr, err := ch.Validator.Struct(payload)
	if err != nil {
		ch.Logger.ErrorT(requestID, "Bad Request", err)
		return web.ResponseErrValidation(ctx, "bad request", mapErr)
	}

	err = ch.UseCaseUsers.StoreUser(userCtx, &payload)
	if err != nil {
		ch.Logger.ErrorT(requestID, "error store data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", "", nil)
}

func (ch *PokemonHandler) UpdateUser(ctx echo.Context) (err error) {
	requestID := mid.GetID(ctx)
	userCtx := mid.SetIDx(ctx.Request().Context(), requestID)

	idP, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ch.Logger.ErrorT(requestID, "error get data", err)
		return web.ResponseFormatter(ctx, http.StatusNotFound, err.Error(), nil, err)
	}

	id := int64(idP)

	var payloads domain.UpdateUserRequest
	if err := ctx.Bind(&payloads); err != nil {
		ch.Logger.ErrorT(requestID, "user update payload", err, mlog.Any("payload", payloads))
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	mapErr, err := ch.Validator.Struct(payloads)
	if err != nil {
		ch.Logger.ErrorT(requestID, "Bad Request", err)
		return web.ResponseErrValidation(ctx, "bad request", mapErr)
	}

	err = ch.UseCaseUsers.UpdateUser(userCtx, id, &payloads)
	if err != nil {
		ch.Logger.ErrorT(requestID, "error update data", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", "", nil)
}

func (ch *PokemonHandler) DeleteUser(c echo.Context) error {
	requestID := mid.GetID(c)
	userCtx := mid.SetIDx(c.Request().Context(), requestID)

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ch.Logger.ErrorT(requestID, "error delete data", err)
		return web.ResponseFormatter(c, http.StatusNotFound, err.Error(), nil, err)
	}

	id := int64(idP)

	err = ch.UseCaseUsers.DeleteUser(userCtx, id)
	if err != nil {
		ch.Logger.ErrorT(requestID, "error delete data", err)
		return web.ResponseFormatter(c, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(c, http.StatusNoContent, "success", "", nil)
}
