package app

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/bestuzheva153/async-task-service/internal/config"
	httpHandler "github.com/bestuzheva153/async-task-service/internal/http/handler"
	"github.com/bestuzheva153/async-task-service/internal/http/router"
	"github.com/bestuzheva153/async-task-service/internal/repository"
	"github.com/bestuzheva153/async-task-service/internal/service"
	"github.com/bestuzheva153/async-task-service/internal/worker"
)

type App struct {
	cfg    *config.Config
	db     *pgxpool.Pool
	router *gin.Engine
}

func New() (*App, error) {
	cfg := config.Load()
	ctx := context.Background()

	db, err := pgxpool.New(ctx, cfg.DBUrl)
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := httpHandler.NewTaskHandler(taskService)

	go worker.NewWorker(taskRepo).Start(ctx)

	r := gin.Default()
	router.Setup(r, taskHandler)

	return &App{
		cfg:    cfg,
		db:     db,
		router: r,
	}, nil
}

func (a *App) Run() error {
	defer a.db.Close()
	return a.router.Run(":8080")
}
