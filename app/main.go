package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ezzycreative1/svc-pokemon/config"
	"github.com/ezzycreative1/svc-pokemon/pkg/db"
	"github.com/ezzycreative1/svc-pokemon/pkg/mlog"
	"github.com/ezzycreative1/svc-pokemon/pkg/mvalidator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type app struct {
	echo      *echo.Echo
	cfg       *config.Group
	logger    mlog.Logger
	database  *gorm.DB
	validator mvalidator.Validator
}

func main() {
	// load config
	cfg := config.LoadConfig()

	// load logger
	logger := mlog.New("info", "stdout")

	// init database
	database := db.NewDatabase(&cfg.Database)
	instDB, _ := database.DB()
	defer instDB.Close()

	// init validator
	mValidator := mvalidator.New()

	// create echo app
	e := echo.New()
	e.HideBanner = true

	// fill app
	application := app{
		echo:      e,
		cfg:       cfg,
		logger:    logger,
		database:  database,
		validator: mValidator,
	}

	// set common middleware
	LoadRoute(&application)

	// Start echo server on goroutine
	go func() {
		e.Server.ReadTimeout = 45 * time.Second
		e.Server.WriteTimeout = 45 * time.Second
		e.Server.IdleTimeout = time.Minute

		if err := e.Start(fmt.Sprintf("0.0.0.0:%v", cfg.Pokemon.HTTPPort)); err != nil && err != http.ErrServerClosed {
			logger.Info("shutting down the server")
			panic(fmt.Sprintf("echo server startup panic: %s", err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// gracefull shutdown stage ===============================================

	logger.Info("shutdown echo server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	// cleanup app ...
	logger.Info("Running cleanup tasks...")

	// close database
	// db, _ := database.()
	// db.Close()

	logger.Info("Done cleanup tasks...")
	logger.Sync()
}
