package workers

import "go-concurrency-sample/internal/types"

type UserRepository interface {
	Insert(user *types.User) error
}
