package execute_worker

import (
	"context"
	"log"
	"sync"

	worker "github.com/contribsys/faktory_worker_go"
	"github.com/tranHieuDev23/IdeTwo/controllers/workers/execute_worker/cpp_job_executor"
	"github.com/tranHieuDev23/IdeTwo/models/daos/source_code_dao"
	"github.com/tranHieuDev23/IdeTwo/models/source_code"
)

type FaktoryWorker struct {
	manager worker.Manager
}

func (worker *FaktoryWorker) Run() {
	worker.manager.Run()
}

var dao = source_code_dao.GetInstance()

func executeJob(ctx context.Context, args ...interface{}) error {
	helper := worker.HelperFor(ctx)
	log.Printf("Working on job %s\n", helper.Jid())

	id := args[0].(string)
	source := dao.GetSourceCode(id)
	if source == nil {
		return nil
	}

	var executor JobExecutor
	switch source.Language {
	case source_code.Cpp:
		executor = cpp_job_executor.GetInstance()
	default:
		panic("Unsupported language")
	}

	output := executor.Execute(*source)

	newSource := source_code.SourceCode{
		Id:       source.Id,
		Language: source.Language,
		Content:  source.Content,
		Input:    source.Input,
		Output:   output,
	}
	dao.UpdateSourceCode(newSource)

	log.Printf("Job %s finished\n", helper.Jid())
	return nil
}

var instance *FaktoryWorker
var once sync.Once

func GetInstance() FaktoryWorker {
	once.Do(func() {
		manager := worker.NewManager()
		manager.Register("Execute Job", executeJob)
		manager.Concurrency = 16
		manager.ProcessStrictPriorityQueues("critical", "default", "bulk")
		instance = &FaktoryWorker{
			manager: *manager,
		}
	})
	return *instance
}
