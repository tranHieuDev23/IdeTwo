package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tranHieuDev23/IdeTwo/controllers/groups/execution_group"
	"github.com/tranHieuDev23/IdeTwo/controllers/groups/source_code_group"
	"github.com/tranHieuDev23/IdeTwo/controllers/workers/execute_worker"
	"github.com/tranHieuDev23/IdeTwo/controllers/workers/execute_worker/cpp_job_executor"
	"github.com/tranHieuDev23/IdeTwo/utils/configs"
)

var conf = configs.GetInstance()

func main() {
	// Create the dedicated execution directory
	if err := os.MkdirAll(conf.IdeTwoExecutionsDir, os.FileMode(0777)); err != nil {
		panic(err)
	}
	// Start Faktory worker in a different goroutine to handle code execution
	worker := execute_worker.GetInstance()
	go worker.Run()
	// Initialize all JobExecutor instances to save time on job handling
	cpp_job_executor.GetInstance()
	// Start HTTP server
	app := gin.Default()
	api := app.Group("/api")
	{
		source_code_group.SourceCodeGroup(*api)
		execution_group.ExecutionGroup(*api)
	}
	app.Run("127.0.0.1:8080")
}
