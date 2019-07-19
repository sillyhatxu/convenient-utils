package rabbitmq

import (
	"github.com/streadway/amqp"
)

var (
	rmqc = New()
)

type QueueConfig struct {
	Arguments  amqp.Table
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Mandatory  bool
	Immediate  bool
	AutoAck    bool
	NoLocal    bool
}

type Config struct {
	URL         string
	QueueConfig QueueConfig
}

func New() *Config {
	return &Config{
		URL: "",
		QueueConfig: QueueConfig{
			Arguments: amqp.Table{
				"name":  "x-message-ttl",
				"value": 3600000,
				"type":  "java.lang.Long",
			},
			Durable:    true,
			AutoDelete: false,
			Exclusive:  false,
			NoWait:     false,
			Mandatory:  false,
			Immediate:  false,
			AutoAck:    true,
			NoLocal:    false,
		},
	}
}

func SetURL(url string) {
	rmqc.URL = url
}

func SetQueueConfig(queueConfig QueueConfig) {
	rmqc.QueueConfig = queueConfig
}
