package rabbitmq

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsumer(t *testing.T) {
	SetURL("amqp://username:password@127.0.0.1:5672/")
	exchange := "exchange.teste"
	routingKey := "routing.key.test" // Key 相当于 kafka topic
	queueValue := "queue.value.test" //相当于  kafka group
	consumer := ConsumerConf{QueueValue: queueValue, Exchange: exchange, RoutingKey: routingKey}
	err := consumer.Consumer(ConsumerMessage{})
	assert.Nil(t, err)
}

type ConsumerMessage struct {
	status string
}

func (cm ConsumerMessage) MessageDelivery(msg amqp.Delivery) {
	var mqGroup MqGroupDTO
	err := json.Unmarshal(msg.Body, &mqGroup)
	if err != nil {
		panic(err)
	}
	log.Info(cm.status)
	log.Info(mqGroup)
}
