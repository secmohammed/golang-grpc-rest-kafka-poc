package types

import (
	"github.com/secmohammed/golang-kafka-grpc-poc/config"
	"github.com/secmohammed/golang-kafka-grpc-poc/pkg/database"
	"github.com/secmohammed/golang-kafka-grpc-poc/pkg/queueing"
)

type Container interface {
	Config() config.Repository
	Database() database.Repository
	Queue() queueing.Messaging
}
