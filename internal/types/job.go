package types 

  type JobType int

const (
    JobInsertUser JobType = iota
    JobSelectUser
)

type Job struct {
    ID      uint64
    Type    JobType
    User    User
}

type Result struct {
    JobID uint64
    Err   error
}

const (
    GenerateUsers JobType = iota
    InsertUsers
    SelectUsers
)
