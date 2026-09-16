package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-concurrency-sample/internal/workers"
)

type Handler struct {
	pool *workers.Pool
}

func NewHandler(pool *workers.Pool) *Handler {
	return &Handler{
		pool: pool,
	}
}

func (h *Handler) GenerateUsers(
	w http.ResponseWriter,
	r *http.Request,
) {
	countStr := r.URL.Query().Get("count")
  fmt.Println("countstr", countStr)

	count, err := strconv.Atoi(countStr)
	if err != nil || count <= 0 {
		http.Error(
			w,
			"count must be a positive integer",
			http.StatusBadRequest,
		)

		return
	}

	jobID, err := h.pool.Submit(
		r.Context(),
		count,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusAccepted)

	_ = json.NewEncoder(w).Encode(map[string]any{
		"job_id": jobID,
		"count":  count,
		"status": "queued",
	})
}
