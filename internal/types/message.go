package types

type MessageType int

const (
    Increment MessageType = iota
    Get
    Ack
)
