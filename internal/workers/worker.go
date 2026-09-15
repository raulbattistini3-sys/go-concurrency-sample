package workers

import (
    "context"
    "log"

    "go-concurrency-sample/internal/database"
    "go-concurrency-sample/internal/types"
)

type Worker struct {
    id   int
    repo *database.UserRepository
}

func NewWorker(
    id int,
    repo *database.UserRepository,
) *Worker {
    return &Worker{
        id:   id,
        repo: repo,
    }
}

func (w *Worker) Run(
    ctx context.Context,
    jobs <-chan types.Job,
    results chan<- types.Result,
) {
    for {
        select {
        case <-ctx.Done():
            return

        case job, ok := <-jobs:
            if !ok {
                return
            }

            err := w.process(job)

            results <- types.Result{
                JobID: job.ID,
                Err:   err,
            }
        }
    }
}

func (w *Worker) process(job types.Job) error {
    log.Printf(
        "[worker:%d] job=%d type=%d",
        w.id,
        job.ID,
        job.Type,
    )

    switch job.Type {
    case types.JobInsertUser:
        return w.repo.Insert(job.User)

    default:
        return nil
    }
}
