package routes

import (
	"github.com/secmohammed/golang-kafka-grpc-poc/types"
	log "github.com/siruspen/logrus"
)

type Type string

const (
	REST Type = "rest"
	GRPC Type = "grpc"
	ALL  Type = "all"
)

// Factory returns the requested config repo
func Factory(t string, c types.Container) Repository {

	switch Type(t) {
	case REST:
		return NewRestRepository(c)
	case GRPC:
		return NewGRPCRepository(c)
	case ALL:
		return NewALLRepository(c)
	default:
		log.Fatalf("Unknown config repository: %s", c)
	}

	return nil
}

// Repository port
type Repository interface {
	Expose() error
}
