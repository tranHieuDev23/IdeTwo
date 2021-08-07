package execute_worker

import (
	"context"
	"log"
	"sync"

	worker "github.com/contribsys/faktory_worker_go"
	"github.com/tranHieuDev23/IdeTwo/controllers/workers/execute_worker/cpp_job_executor"
	"github.com/tranHieuDev23/IdeTwo/controllers/workers/execute_worker/job_executor"
	"github.com/tranHieuDev23/IdeTwo/models/daos/execution_dao"
	"github.com/tranHieuDev23/IdeTwo/models/daos/source_code_dao"
	"github.com/tranHieuDev23/IdeTwo/models/source_code"
	"github.com/tranHieuDev23/IdeTwo/utils/configs"
)

type FaktoryWorker struct {
	manager worker.Manager
}

func (worker *FaktoryWorker) Run() {
	worker.manager.Run()
}

var sourceDao = source_code_dao.GetInstance()
var executionDao = execution_dao.GetInstance()

func executeJob(ctx context.Context, args ...interface{}) error {
	helper := worker.HelperFor(ctx)
	log.Printf("Working on job %s\n", helper.Jid())

	id := args[0].(string)
	exec := executionDao.GetExecution(id)
	if exec == nil {
		return nil
	}
	source := sourceDao.GetSourceCode(exec.OfSourceCodeId)
	if source == nil {
		return nil
	}

	var executor job_executor.JobExecutor
	switch source.Language {
	case source_code.Cpp:
		executor = cpp_job_executor.GetInstance()
	default:
		panic("Unsupported language")
	}

	output := executor.Execute(*source)

	exec.Status = output.Status
	exec.RunTime = output.RunTime
	exec.Output = output.Output
	executionDao.UpdateExecution(*exec)

	log.Printf("Job %s finished\n", helper.Jid())
	return nil
}

var instance *FaktoryWorker
var once sync.Once
var conf = configs.GetInstance()

func GetInstance() FaktoryWorker {
	once.Do(func() {
		manager := worker.NewManager()
		manager.Register("Execute Job", executeJob)
		manager.Concurrency = conf.FaktoryWorkerConcurrency
		manager.ProcessStrictPriorityQueues("critical", "default", "bulk")
		instance = &FaktoryWorker{
			manager: *manager,
		}
	})
	return *instance
}
