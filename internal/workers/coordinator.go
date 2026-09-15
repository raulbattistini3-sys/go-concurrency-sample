package workers

import (
    "context"

    "go-concurrency-sample/internal/users"
    "go-concurrency-sample/internal/types"
)

type Coordinator struct {
    jobs    chan<- types.Job
    results <-chan types.Result
}

func NewCoordinator(
    jobs chan<- types.Job,
    results <-chan types.Result,
) *Coordinator {
    return &Coordinator{
        jobs:    jobs,
        results: results,
    }
}

func (c *Coordinator) Dispatch(ctx context.Context) error {
    for i := 0; i < 10_000; i++ {
        user := users.Generate(i)

        job := types.Job{
            ID:   uint64(i),
            Type: types.JobInsertUser,
            User: user,
        }

        select {
        case <-ctx.Done():
            return ctx.Err()

        case c.jobs <- job:
        }
    }

    return nil
}
