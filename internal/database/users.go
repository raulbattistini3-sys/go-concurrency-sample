package database

import (
    "database/sql"

    "github.com/raulbattistini/parallel-go/internal/types"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{
        db: db,
    }
}

func (r *UserRepository) Insert(user types.User) error {
    _, err := r.db.Exec(`
        INSERT INTO users (
            cpf,
            name,
            email,
            birthdate,
            phone,
            mother_name,
            address,
            address_number,
            address_complement,
            neighborhood,
            city,
            state,
            zipcode,
            occupation,
            marital_status,
            gender
        )
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `,
        user.CPF,
        user.Name,
        user.Email,
        user.Birthdate,
        user.Phone,
        user.MotherName,
        user.Address,
        user.AddressNumber,
        user.AddressComplement,
        user.Neighborhood,
        user.City,
        user.State,
        user.Zipcode,
        user.Occupation,
        user.MaritalStatus,
        user.Gender,
    )

    return err
}
