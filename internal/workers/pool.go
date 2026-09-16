package workers

import (
	"context"
	"sync"
	"sync/atomic"

	"go-concurrency-sample/internal/database"
	"go-concurrency-sample/internal/types"
	"go-concurrency-sample/internal/users"
)

type Pool struct {
	repo        *database.UserRepository
	workerCount int

	Jobs    chan types.Job
	Results chan types.Result

	jobID atomic.Uint64
}

func NewPool(
	repo *database.UserRepository,
	workerCount int,
	queueSize int,
) *Pool {
	return &Pool{
		repo:        repo,
		workerCount: workerCount,
		Jobs:         make(chan types.Job, queueSize),
		Results:      make(chan types.Result, queueSize),
	}
}

func (p *Pool) Submit(
	ctx context.Context,
	count int,
) (uint64, error) {
	workloadID := p.jobID.Add(1)

	for i := 0; i < count; i++ {
		job := types.Job{
			ID:   p.jobID.Add(1),
			Type: types.JobInsertUser,
			User: users.Generate(i),
		}

		select {
		case <-ctx.Done():
			return workloadID, ctx.Err()

		case p.Jobs <- job:
		}
	}

	return workloadID, nil
}

func (p *Pool) Run(ctx context.Context) {
	var wg sync.WaitGroup

	for i := 0; i < p.workerCount; i++ {
		worker := NewWorker(i, p.repo)

		wg.Add(1)

		go func() {
			defer wg.Done()

			worker.Run(
				ctx,
				p.Jobs,
				p.Results,
			)
		}()
	}

	wg.Wait()
}
