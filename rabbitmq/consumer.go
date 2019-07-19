package rabbitmq

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type ConsumerConf struct {
	QueueValue string // @Queue(value='') 相当于  kafka group

	Exchange string

	RoutingKey string // Key 相当于 kafka topic
}

func (cc ConsumerConf) String() string {
	return fmt.Sprintf("{ QueueValue : %s, Exchange : %s, RoutingKey : %s }", cc.QueueValue, cc.Exchange, cc.RoutingKey)
}

type ConsumerInterface interface {
	MessageDelivery(msg amqp.Delivery)
}

func (cc ConsumerConf) Consumer(ci ConsumerInterface) error {
	log.Infof("RabbitMQ ConsumerConf : %v", cc)
	conn, err := amqp.Dial(rmqc.URL)
	if err != nil {
		log.Error("Connection RabbitMQ error.", err)
		return err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Error("Get RabbitMQ channel error.", err)
		return err
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		cc.RoutingKey,               // name
		rmqc.QueueConfig.Durable,    // durable
		rmqc.QueueConfig.AutoDelete, // delete when unused
		rmqc.QueueConfig.Exclusive,  // exclusive
		rmqc.QueueConfig.NoWait,     // no-wait
		rmqc.QueueConfig.Arguments,  // arguments
	)
	if err != nil {
		log.Error("Get RabbitMQ queue error.", err)
		return err
	}
	err = ch.QueueBind(
		q.Name,        // queue name
		cc.RoutingKey, // routing key
		cc.Exchange,   // exchange
		rmqc.QueueConfig.NoWait,
		rmqc.QueueConfig.Arguments)
	if err != nil {
		log.Error("RabbitMQ set bind error.", err)
		return err
	}
	msgs, err := ch.Consume(
		q.Name,                     // queue
		"",                         // consumer
		rmqc.QueueConfig.AutoAck,   // auto-ack
		rmqc.QueueConfig.Exclusive, // exclusive
		rmqc.QueueConfig.NoLocal,   // no-local
		rmqc.QueueConfig.NoWait,    // no-wait
		rmqc.QueueConfig.Arguments, // args
	)
	if err != nil {
		log.Error("RabbitMQ consume error.", err)
		return nil
	}
	forever := make(chan bool)
	go func() {
		for delivery := range msgs {
			ci.MessageDelivery(delivery)
		}
	}()
	log.Info("Waiting for messages.")
	<-forever
	log.Warningf("RabbitMQ Consumer Exit; ConsumerConf : %v", cc)
	return nil
}
