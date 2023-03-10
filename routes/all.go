package routes

import (
    "github.com/secmohammed/golang-kafka-grpc-poc/types"
)

type all struct {
    g *grpcClient
    r *rest
    q *messagingQueueHandler
}

func NewALLRepository(c types.Container) *all {
    return &all{g: NewGRPCRepository(c), r: NewRestRepository(c), q: NewMessagingQueueRepository(c)}
}
func (a *all) Expose() error {
    ch := make(chan error, 3)
    go func() {
        if err := a.g.Expose(); err != nil {
            ch <- err
            close(ch)
        }
    }()
    go func() {
        if err := a.r.Expose(); err != nil {
            ch <- err
            close(ch)
        }
    }()
    go func() {
        if err := a.q.Expose(); err != nil {
            ch <- err
            close(ch)
        }
    }()
    for err := range ch {
        return err
    }
    return nil
}
