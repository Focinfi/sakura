package worker

import (
	"time"

	"gopkg.in/go-playground/pool.v3"
)

// OnSuccessFunc handles work completes successfully
type OnSuccessFunc func(result interface{})

// OnErrorFunc handles work errors
type OnErrorFunc func(err error)

// Worker for worker
type Worker struct {
	Number    uint
	Retry     bool
	Queue     string
	Work      pool.WorkFunc
	OnSuccess OnSuccessFunc
	OnError   OnErrorFunc
}

// PerformAsync performs immediately
func (worker *Worker) PerformAsync() {
	p := pool.NewLimited(worker.Number)
	defer p.Close()
DO:
	work := p.Queue(worker.Work)
	work.Wait()
	if err := work.Error(); err != nil && worker.Retry {
		worker.Retry = false
		goto DO
	}

	if err := work.Error(); err != nil {
		worker.OnError(err)
	} else {
		worker.OnSuccess(work.Value())
	}
}

// PerformIn performs work in the given second
func (worker Worker) PerformIn(second int) {
	time.AfterFunc(time.Second*time.Duration(second), func() {
		worker.PerformAsync()
	})
}

// PerformAt performs work at the given moment
func (worker Worker) PerformAt(at time.Time) {
	worker.PerformIn(int(at.Unix() - time.Now().Unix()))
}

// New allocates and returns a new Worker
func New(num uint) *Worker {
	return &Worker{
		Number:    num,
		Queue:     "default",
		Work:      func(wu pool.WorkUnit) (interface{}, error) { return nil, nil },
		OnSuccess: func(result interface{}) {},
		OnError:   func(err error) {},
	}
}

// SetRetry set retry
func (worker *Worker) SetRetry(retry bool) *Worker {
	worker.Retry = retry
	return worker
}

// SetQueue set queue
func (worker *Worker) SetQueue(queue string) *Worker {
	worker.Queue = queue
	return worker
}

// SetWork set work
func (worker *Worker) SetWork(work pool.WorkFunc) *Worker {
	worker.Work = work
	return worker
}

// SetOnSuccess set OnSuccess for worker
func (worker *Worker) SetOnSuccess(onSuccess OnSuccessFunc) *Worker {
	worker.OnSuccess = onSuccess
	return worker
}

// SetOnError set OnError for worker
func (worker *Worker) SetOnError(onError OnErrorFunc) *Worker {
	worker.OnError = onError
	return worker
}
