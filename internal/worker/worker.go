package worker

import (
	"context"
	"log"
	"math/rand"
	"time"

	"awesomeProject1/internal/model"
	"awesomeProject1/internal/repository"
)

type Worker struct {
	repo *repository.TaskRepository
}

func NewWorker(repo *repository.TaskRepository) *Worker {
	return &Worker{repo: repo}
}

func (w *Worker) Start(ctx context.Context) {
	for {
		task, err := w.repo.FetchPending(ctx)
		if err != nil {
			log.Println("fetch error:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		if task == nil {
			time.Sleep(2 * time.Second)
			continue
		}

		log.Println("processing task:", task.ID)

		time.Sleep(time.Duration(rand.Intn(3)+2) * time.Second)

		if rand.Intn(10) < 2 {
			errMsg := "random failure"
			w.repo.UpdateStatus(ctx, task.ID, model.StatusFailed, nil, &errMsg)
			continue
		}

		result := "success"
		w.repo.UpdateStatus(ctx, task.ID, model.StatusDone, &result, nil)
	}
}
