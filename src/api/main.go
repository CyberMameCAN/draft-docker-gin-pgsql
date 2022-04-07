package main

import (
	"gopgsql/db"
	"log"

	TaskController "gopgsql/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server..")

	db.Init()
	// サンプルデータの投入
	db.InsertRecord()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.GET("/", TaskController.GetTasks) // /api/v1/tasks
			tasks.POST("/", TaskController.CreateTask)
			tasks.PUT("/:id", TaskController.UpdateTask)
			tasks.DELETE("/:id", TaskController.DeleteTask)
		}
	}

	err := r.Run(":7070")
	if err != nil {
		log.Fatal("サーバ起動に失敗", err)
	}
}
