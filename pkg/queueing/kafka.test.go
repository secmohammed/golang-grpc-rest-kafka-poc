package queueing

type kafkaTestConn struct {
}

func (k kafkaTestConn) Read(topic string, callbacks []CallbackFunc) error {
	return nil
}

func (k kafkaTestConn) Write(topic string, key, value []byte) error {
	return nil
}

func (k kafkaTestConn) Close() error {
	return nil
}

func NewTest() (Messaging, error) {
	return &kafkaTestConn{}, nil
}
