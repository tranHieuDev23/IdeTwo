package faktory_proxy

import (
	"sync"

	faktory "github.com/contribsys/faktory/client"
	"github.com/tranHieuDev23/IdeTwo/utils/configs"
)

// Abstract proxy to communicate with Faktory server.
type FaktoryProxy struct {
	pool faktory.Pool
}

var instance *FaktoryProxy
var once sync.Once

func GetInstance() FaktoryProxy {
	once.Do(func() {
		c := configs.GetInstance()
		pool, err := faktory.NewPool(c.FaktoryPoolCapacity)
		if err != nil {
			panic(err)
		}
		instance = &FaktoryProxy{pool: *pool}
	})
	return *instance
}

// Push a new Execution Job to Faktory, to execute the source code with the
// provided id and update the output.
func (proxy *FaktoryProxy) PushExecuteJob(id string) {
	client, err := proxy.pool.Get()
	if err != nil {
		panic(err)
	}
	defer proxy.pool.Put(client)
	job := faktory.NewJob("Execute Job", id)
	err = client.Push(job)
	if err != nil {
		panic(err)
	}
}
