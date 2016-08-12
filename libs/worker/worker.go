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
	Pool      pool.Pool
	Retry     bool
	Works     []pool.WorkFunc
	OnSuccess OnSuccessFunc
	OnError   OnErrorFunc
}

// PerformAsync performs immediately
func (worker *Worker) PerformAsync(work pool.WorkFunc) {
	if worker.Pool == nil {
		worker.Pool = pool.NewLimited(10)
		defer worker.Pool.Close()
	}

	worker.Works = append(worker.Works, work)
	task := worker.Pool.Queue(work)
	task.Wait()
	if err := task.Error(); err != nil && worker.Retry {
		again := worker.Pool.Queue(work)
		again.Wait()
	}

	if err := task.Error(); err != nil && worker.OnSuccess != nil {
		worker.OnError(err)
	} else if worker.OnError != nil {
		worker.OnSuccess(task.Value())
	}
}

// PerformIn performs work in the given second
func (worker Worker) PerformIn(second int, work pool.WorkFunc) {
	time.AfterFunc(time.Second*time.Duration(second), func() {
		worker.PerformAsync(work)
	})
}

// PerformAt performs work at the given moment
func (worker Worker) PerformAt(at time.Time, work pool.WorkFunc) {
	worker.PerformIn(int(at.Unix()-time.Now().Unix()), work)
}

// New allocates and returns a new Worker
func New(p pool.Pool) *Worker {
	return &Worker{
		Pool:      p,
		Works:     []pool.WorkFunc{},
		OnSuccess: func(result interface{}) {},
		OnError:   func(err error) {},
	}
}

// SetRetry set retry
func (worker *Worker) SetRetry(retry bool) *Worker {
	worker.Retry = retry
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
