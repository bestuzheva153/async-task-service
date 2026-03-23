package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"awesomeProject1/internal/config"
	"awesomeProject1/internal/handler"
	"awesomeProject1/internal/repository"
	"awesomeProject1/internal/service"
	"awesomeProject1/internal/worker"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()

	db, err := pgxpool.New(ctx, cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTaskRepository(db)
	service := service.NewTaskService(repo)
	handler := handler.NewTaskHandler(service)

	go worker.NewWorker(repo).Start(ctx)

	r := gin.Default()

	r.POST("/tasks", handler.CreateTask)
	r.GET("/tasks/:id", handler.GetTask)
	r.GET("/tasks", handler.GetAll)

	log.Println("server started on :8080")
	r.Run(":8080")
}
