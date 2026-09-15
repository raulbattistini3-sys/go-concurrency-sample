package users

import (
	"fmt"
	"time"

	"go-concurrency-sample/internal/types"
)

func Generate(index int) types.User { 
  return types.User{ 
    CPF: fmt.Sprintf("%011d", index), 
    Name: fmt.Sprintf("User %d", index), 
    Email: fmt.Sprintf("user-%d@example.com", index), 
    Birthdate: time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), 
  }
}
