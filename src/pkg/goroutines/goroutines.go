package goroutines

import (
	"runtime"
	"sync"
)

var (
	jobsNum    = 1024 * runtime.NumCPU()
	workersNum = runtime.NumCPU()
)

type Context struct {
	Jobs      chan func()
	WaitGroup sync.WaitGroup
}

func (context *Context) Add(job func()) {
	context.Jobs <- job
	context.WaitGroup.Add(1)
}

func (context *Context) Done() {
	context.WaitGroup.Done()
}

func (context *Context) Wait() {
	close(context.Jobs)
	context.WaitGroup.Wait()
}

func New() *Context {
	context := &Context{
		Jobs: make(chan func(), jobsNum),
	}

	for range workersNum {
		go func() {
			for job := range context.Jobs {
				job()
				context.Done()
			}
		}()
	}

	return context
}
