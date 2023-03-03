package queueing

import (
    "context"
    "errors"
    "fmt"
    "github.com/segmentio/kafka-go"
    "github.com/segmentio/kafka-go/compress"
    _ "github.com/segmentio/kafka-go/gzip"
    _ "github.com/segmentio/kafka-go/snappy"
    "github.com/siruspen/logrus"
    "strings"
    "sync"
    "time"
)

type CallbackFunc func(message kafka.Message) error
type Compression string
type Messaging interface {
    Read(topic string, callbacks []CallbackFunc) error
    Write(topic string, key, value []byte) error
    Close() error
}
type Option struct {
    Host              []string
    ConsumerGroup     string
    Interval          int
    RequiredAck       int
    QueueCapacity     int
    MinBytes          int
    MaxBytes          int
    HeartbeatInterval time.Duration
    ReadBackoffMin    time.Duration
    ReadBackoffMax    time.Duration
    CommitInterval    time.Duration
    CompressionCodec  Compression
}

const (
    Snappy = "snappy"
    Gzip   = "gzip"
)

type kafkaClient struct {
    option  Option
    readers map[string]*kafka.Reader
    writers map[string]*kafka.Writer
    ctx     context.Context
    log     *logrus.Logger
    mu      *sync.Mutex
}

func getOption(option *Option) error {
    if len(option.Host) == 0 {
        return errors.New("host is required")
    }
    if option.ConsumerGroup == "" {
        return errors.New("consumerGroup is required")
    }
    if option.Interval == 0 {
        option.Interval = 1
    }
    if option.RequiredAck == 0 {
        option.RequiredAck = -1
    }
    if option.QueueCapacity <= 0 {
        option.QueueCapacity = 100
    }
    if option.HeartbeatInterval <= 0 {
        option.HeartbeatInterval = 3 * time.Second
    }
    if option.ReadBackoffMin <= 0 {
        option.ReadBackoffMin = 100 * time.Millisecond
    }
    if option.ReadBackoffMax <= 0 {
        option.ReadBackoffMax = 1 * time.Second
    }

    if option.CompressionCodec == "" {
        option.CompressionCodec = Gzip
    }
    if option.CompressionCodec != Snappy && option.CompressionCodec != Gzip {
        return errors.New("error compression codec type")
    }
    return nil
}

func (k *kafkaClient) Read(topic string, callbacks []CallbackFunc) error {
    if len(callbacks) < 1 {
        return errors.New("at least 1 callbacks is required")
    }

    k.mu.Lock()
    if _, ok := k.readers[topic]; !ok {
        reader := kafka.NewReader(kafka.ReaderConfig{
            Brokers:           k.option.Host,
            GroupID:           k.option.ConsumerGroup,
            Topic:             topic,
            MaxWait:           time.Duration(k.option.Interval) * time.Millisecond,
            QueueCapacity:     k.option.QueueCapacity,
            HeartbeatInterval: k.option.HeartbeatInterval,
            ReadBackoffMin:    k.option.ReadBackoffMin,
            ReadBackoffMax:    k.option.ReadBackoffMax,
            CommitInterval:    k.option.CommitInterval,
            MinBytes:          k.option.MinBytes,
            MaxBytes:          k.option.MaxBytes,
        })
        k.readers[topic] = reader
    }
    k.mu.Unlock()

    reader := k.readers[topic]
    for {
        m, err := reader.ReadMessage(k.ctx)
        if err != nil {
            k.log.Error(err)
            continue
        }

        for _, c := range callbacks {
            if err = c(m); err != nil {
                k.log.Error(err)
            }
        }
    }
}

func (k *kafkaClient) Close() error {
    var err error
    // - close writer
    for _, w := range k.writers {
        if e := w.Close(); e != nil {
            err = e
            k.log.Error(err)
        }
    }

    // - close reader
    for _, r := range k.readers {
        if e := r.Close(); e != nil {
            err = e
            k.log.Error(err)
        }
    }

    return err
}
func (k *kafkaClient) Write(topic string, key, value []byte) error {
    k.mu.Lock()

    var compressionCodec kafka.Compression
    switch k.option.CompressionCodec {
    case Snappy:
        compressionCodec = compress.Snappy
    case Gzip:
        compressionCodec = compress.Gzip
    default:
        k.mu.Unlock()
        return errors.New("error compression codec")
    }

    if _, ok := k.writers[topic]; !ok {
        writer := &kafka.Writer{
            Addr:                   kafka.TCP(strings.Join(k.option.Host, ",")),
            Topic:                  topic,
            Balancer:               &kafka.Hash{},
            RequiredAcks:           kafka.RequiredAcks(k.option.RequiredAck),
            BatchTimeout:           time.Duration(k.option.Interval) * time.Millisecond,
            Compression:            compressionCodec,
            AllowAutoTopicCreation: true,
        }
        k.writers[topic] = writer
    }
    k.mu.Unlock()

    w := k.writers[topic]
    if err := w.WriteMessages(context.Background(), kafka.Message{Value: value, Key: key}); err != nil {
        k.log.Error(err)
        return fmt.Errorf(err.Error(), "failed to publish message on topic %s", topic)
    }
    return nil
}

func New(option Option, ctx context.Context) (Messaging, error) {
    err := getOption(&option)
    if err != nil {
        return nil, err
    }
    return &kafkaClient{
        option:  option,
        mu:      &sync.Mutex{},
        log:     logrus.New(),
        ctx:     ctx,
        writers: make(map[string]*kafka.Writer),
        readers: make(map[string]*kafka.Reader),
    }, nil
}
