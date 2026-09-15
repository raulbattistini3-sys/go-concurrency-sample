package users

import (
    "fmt"

    "internal/types"
)

func Generate(index int) types.User {
    return types.User{
        CPF:   generateCPF(),
        Name:  fmt.Sprintf("User %d", index),
        Email: fmt.Sprintf("user-%d@example.com", index),
    }
}
