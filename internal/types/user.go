package types

import "time"

type User struct {
    CPF              string
    Name             string
    Email            string
    Birthdate        time.Time
    Phone            *string
    MotherName       *string
    Address          *string
    AddressNumber    *string
    AddressComplement *string
    Neighborhood     *string
    City             *string
    State            *string
    Zipcode          *string
    Occupation       *string
    MaritalStatus    *string
    Gender           *string
}
