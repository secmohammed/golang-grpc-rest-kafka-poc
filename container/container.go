package container

import (
	"github.com/secmohammed/golang-kafka-grpc-poc/config"
	"github.com/secmohammed/golang-kafka-grpc-poc/pkg/database"
	"github.com/secmohammed/golang-kafka-grpc-poc/types"

	"sync"
)

var (
	instantiateAppOnce sync.Once
	appInstance        *container
)

type container struct {
	c  config.Repository
	db database.Repository
}

func NewApplication(c config.Repository) types.Container {
	instantiateAppOnce.Do(func() {

		appInstance = &container{
			c:  c,
			db: database.NewDatabaseConnection(c),
		}
	})
	return appInstance
}

func (c *container) Get() *container {
	return c
}
func (c *container) Database() database.Repository {
	return c.db
}
func (c *container) Config() config.Repository {
	return c.c
}
