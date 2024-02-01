package parkingattendant

type ObserverEvent int

const (
    FULL ObserverEvent = iota
    EMPTY
)

type Statergy int

const (
    NEAREST Statergy = iota
    FARTHEST
)
