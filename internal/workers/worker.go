package workers

import (
	"context"
	"log"

	"go-concurrency-sample/internal/types"
)

type Worker struct {
	id   int
	repo UserRepository
}

func NewWorker(
	id int,
	repo UserRepository,
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
      if err != nil {
        log.Printf("error when running worker: %v", err)
      }

			results <- types.Result{
				JobID: job.ID,
			}
		}
	}
}

func (w *Worker) process(job types.Job) error {
	switch job.Type {
	case types.JobInsertUser:
		return w.repo.Insert(&job.User)

	default:
		log.Printf(
			"[worker:%d] unknown job type=%d",
			w.id,
			job.Type,
		)

		return nil
	}
}

