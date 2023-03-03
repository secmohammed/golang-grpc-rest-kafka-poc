package container

import (
    "context"
    "fmt"
    "github.com/secmohammed/golang-kafka-grpc-poc/config"
    "github.com/secmohammed/golang-kafka-grpc-poc/pkg/database"
    "github.com/secmohammed/golang-kafka-grpc-poc/pkg/queueing"
    "github.com/secmohammed/golang-kafka-grpc-poc/types"
    "log"

    "sync"
)

var (
    instantiateAppOnce sync.Once
    appInstance        *container
)

type container struct {
    c  config.Repository
    db database.Repository
    q  queueing.Messaging
}

func NewApplication(c config.Repository) types.Container {
    instantiateAppOnce.Do(func() {

        env, err := c.GetString("app.env")
        if err != nil {
            log.Fatal(err)
        }
        var q queueing.Messaging
        if env == "test" {
            q, _ = queueing.NewTest()

        } else {
            messagingHosts, err := c.GetStringSlice("app.messaging.hosts")
            if err != nil {
                log.Fatal(err)
            }
            messagingPort, err := c.GetString("app.messaging.port")
            if err != nil {
                log.Fatal(err)
            }
            var hosts []string
            for _, host := range messagingHosts {
                hosts = append(hosts, fmt.Sprintf("%s:%s", host, messagingPort))
            }
            q, err = queueing.New(queueing.Option{
                Host:          hosts,
                ConsumerGroup: "test-consumer",
            }, context.Background())

            if err != nil {
                log.Fatal(err)
            }
        }

        appInstance = &container{
            c:  c,
            q:  q,
            db: database.NewDatabaseConnection(c),
        }
    })
    return appInstance
}

func (c *container) Get() *container {
    return c
}
func (c *container) Queue() queueing.Messaging {
    return c.q
}
func (c *container) Database() database.Repository {
    return c.db
}
func (c *container) Config() config.Repository {
    return c.c
}
