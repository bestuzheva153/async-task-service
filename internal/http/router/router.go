package router

import (
	"github.com/gin-gonic/gin"

	"github.com/bestuzheva153/async-task-service/internal/http/handler"
)

func Setup(r *gin.Engine, taskHandler *handler.TaskHandler) {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			tasks := v1.Group("/tasks")
			{
				tasks.POST("", taskHandler.CreateTask)
				tasks.GET("", taskHandler.GetAll)
				tasks.GET("/:id", taskHandler.GetTask)
			}
		}
	}
}
