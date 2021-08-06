package configs

import (
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

// Contain configuration settings extracted from .env file.
type IdeTwoConfigs struct {
	IdeTwoExecutionsDir      string
	MongoDbUri               string
	MongoDbDb                string
	FaktoryPoolCapacity      int
	FaktoryWorkerConcurrency int
}

var instance *IdeTwoConfigs = nil
var once sync.Once

func GetInstance() IdeTwoConfigs {
	once.Do(func() {
		env, err := godotenv.Read()
		if err != nil {
			panic(err)
		}
		faktoryPoolCapacity, err := strconv.Atoi(env["FAKTORY_POOL_CAPACITY"])
		if err != nil {
			panic(err)
		}
		faktoryWorkerConcurrency, err := strconv.Atoi(env["FAKTORY_WORKER_CONCURRENCY"])
		if err != nil {
			panic(err)
		}
		instance = &IdeTwoConfigs{
			IdeTwoExecutionsDir:      env["IDETWO_EXECUTIONS_DIR"],
			MongoDbUri:               env["MONGODB_URI"],
			MongoDbDb:                env["MONGODB_DB"],
			FaktoryPoolCapacity:      faktoryPoolCapacity,
			FaktoryWorkerConcurrency: faktoryWorkerConcurrency,
		}
	})
	return *instance
}
