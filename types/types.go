package types

import (
	"github.com/secmohammed/golang-kafka-grpc-poc/config"
	"github.com/secmohammed/golang-kafka-grpc-poc/pkg/database"
)

type Container interface {
	Config() config.Repository
	Database() database.Repository
}
