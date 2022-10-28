package main

import (
	"net/http"

	"github.com/ezzycreative1/svc-pokemon/app/v1/handler"
	"github.com/ezzycreative1/svc-pokemon/pkg/web"

	"github.com/ezzycreative1/svc-pokemon/internal/core/usecase"
	"github.com/ezzycreative1/svc-pokemon/internal/repository/mysql"
	"github.com/labstack/echo/v4"
)

func LoadRoute(app *app) {
	// init dependency
	roleRepos := mysql.NewMysqlRolesRepo(app.database)
	userRepos := mysql.NewMysqlUserRepo(app.database)
	pokeRepos := mysql.NewMysqlPokemonRepo(app.database)
	roleUCase := usecase.NewRolesUsecase(
		&roleRepos,
	)
	userUCase := usecase.NewUserUsecase(
		&userRepos,
	)
	pokemonUCase := usecase.NewPokemonUsecase(
		&pokeRepos,
	)
	// create handler
	pokemonHandler := handler.NewPokemonHandler(
		&roleUCase,
		&userUCase,
		&pokemonUCase,
		app.validator,
		app.logger,
		*app.cfg,
	)

	// init additional middleware here or directly in route (ex. JWT, api key)
	// ...

	// set route =============================================================
	// route for check health
	app.echo.GET("v1/health/ping", func(c echo.Context) error {
		return web.ResponseFormatter(c, http.StatusOK, "Success", map[string]any{"status": "ok"}, nil)
	})

	g := app.echo.Group("/v1/pokemon")
	//router role
	g.POST("/role", pokemonHandler.StoreRole)
	g.GET("/roles", pokemonHandler.FetchRoles)
	g.PUT("/role/:id", pokemonHandler.UpdateRole)
	g.GET("/role/:id", pokemonHandler.GetRoleByID)
	g.DELETE("/role/:id", pokemonHandler.DeleteRole)
	//router user
	g.POST("/user", pokemonHandler.StoreUser)
	g.GET("/users", pokemonHandler.FetchUsers)
	g.PUT("/user/:id", pokemonHandler.UpdateUser)
	g.GET("/user/:id", pokemonHandler.GetUserByID)
	g.DELETE("/user/:id", pokemonHandler.DeleteUser)
	g.POST("/user/login", pokemonHandler.Login)

	g.GET("/list", pokemonHandler.FetchPokemons)
	g.POST("/create", pokemonHandler.StorePokemon)
}
