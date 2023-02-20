package routes

import (
	"fmt"
	"github.com/secmohammed/golang-kafka-grpc-poc/pkg/queueing"
	"github.com/secmohammed/golang-kafka-grpc-poc/types"
	kafka2 "github.com/segmentio/kafka-go"
)

type TopicHandler map[string]func(message kafka2.Message) error
type QueueHandler interface {
	RegisterReaders() error
}
type messagingQueueHandler struct {
	c             types.Container
	topicHandlers map[string]TopicHandler
}

func (q *messagingQueueHandler) RegisterReaders() error {
	errCh := make(chan error)

	go func() {
		if err := q.c.Queue().Read("companies", []queueing.CallbackFunc{
			func(message kafka2.Message) error {
				if _, ok := q.topicHandlers["companies"][string(message.Key)]; !ok {
					return fmt.Errorf("call to undefined handler: %s for topic: %s", string(message.Key), message.Topic)
				}
				return q.topicHandlers["companies"][string(message.Key)](message)
			},
		}); err != nil {
			errCh <- err
			close(errCh)
		}
	}()
	go func() {
		if err := q.c.Queue().Read("users", []queueing.CallbackFunc{
			func(message kafka2.Message) error {
				if _, ok := q.topicHandlers["users"][string(message.Key)]; !ok {
					return fmt.Errorf("call to undefined handler: %s for topic: %s", string(message.Key), message.Topic)
				}
				return q.topicHandlers["users"][string(message.Key)](message)
			},
		}); err != nil {
			errCh <- err
			close(errCh)
		}
	}()
	return <-errCh
}
func registerCompanyHandlers(c types.Container) TopicHandler {
	// instantiate use-cases/company-related handlers
	return TopicHandler{
		"getAll": func(message kafka2.Message) error {
			env, _ := c.Config().GetString("app.env")
			fmt.Println(string(message.Key), string(message.Value), env)
			return nil
		},
		"getOne": func(message kafka2.Message) error {
			env, _ := c.Config().GetString("app.env")
			fmt.Println(string(message.Key), string(message.Value), env)
			return nil
		},
		"create": func(message kafka2.Message) error {
			env, _ := c.Config().GetString("app.env")
			fmt.Println(string(message.Key), string(message.Value), env)
			return nil
		},
		"update": func(message kafka2.Message) error {
			env, _ := c.Config().GetString("app.env")
			fmt.Println(string(message.Key), string(message.Value), env)
			return nil
		},
		"delete": func(message kafka2.Message) error {
			env, _ := c.Config().GetString("app.env")
			fmt.Println(string(message.Key), string(message.Value), env)
			return nil
		},
	}
}
func registerUserHandlers(c types.Container) TopicHandler {
	return TopicHandler{
		"login": func(message kafka2.Message) error {
			env, _ := c.Config().GetString("app.env")
			fmt.Println(string(message.Key), string(message.Value), env)
			return nil
		},
		"register": func(message kafka2.Message) error {
			env, _ := c.Config().GetString("app.env")
			fmt.Println(string(message.Key), string(message.Value), env)
			return nil
		},
	}
}
func NewMessagingQueueRepository(c types.Container) *messagingQueueHandler {
	topicHandlers := make(map[string]TopicHandler)
	topicHandlers["companies"] = registerCompanyHandlers(c)
	topicHandlers["users"] = registerUserHandlers(c)
	return &messagingQueueHandler{c: c, topicHandlers: topicHandlers}
}
func (q *messagingQueueHandler) Expose() error {
	if err := q.RegisterReaders(); err != nil {
		if err2 := q.c.Queue().Close(); err2 != nil {
			return err2
		}
		return err
	}
	return nil
}
