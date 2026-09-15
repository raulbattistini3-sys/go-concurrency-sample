package workers

import (
      "context"
    "sync"

    "go-concurrency-sample/internal/database"
    "go-concurrency-sample/internal/types"
)

type Pool struct {
    repo        *database.UserRepository
    workerCount int
}

func NewPool(
    repo *database.UserRepository,
    workerCount int,
) *Pool {
    return &Pool{
        repo:        repo,
        workerCount: workerCount,
    }
}

func (p *Pool) Run(ctx context.Context) error {
    jobs := make(chan types.Job, 1000)
    results := make(chan types.Result, 1000)

    var wg sync.WaitGroup

    for i := 0; i < p.workerCount; i++ {
        worker := NewWorker(i, p.repo)

        wg.Add(1)

        go func() {
            defer wg.Done()

            worker.Run(ctx, jobs, results)
        }()
    }

    // coordinator
    coordinator := NewCoordinator(jobs, results)

    if err := coordinator.Dispatch(ctx); err != nil {
        return err
    }

    close(jobs)

    wg.Wait()

    return nil
}
