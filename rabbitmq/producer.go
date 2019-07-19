package rabbitmq

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type ProducerConf struct {
	Exchange string

	RoutingKey string // Key 相当于 kafka topic
}

func (pc ProducerConf) String() string {
	return fmt.Sprintf("{ Exchange : %s, RoutingKey : %s }", pc.Exchange, pc.RoutingKey)
}

func (pc ProducerConf) Send(producer interface{}) error {
	log.Infof("RabbitMQ ProducerConf : %v", pc)
	producerJSON, err := json.Marshal(producer)
	if err != nil {
		return err
	}
	log.Info(string(producerJSON))
	if len(producerJSON) <= 2 {
		//JSON is "{}"
		return errors.New("Struct to json error.")
	}
	conn, err := amqp.Dial(rmqc.URL)
	if err != nil {
		log.Errorf("Connection [%v] RabbitMQ error.", rmqc.URL, err)
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
		pc.RoutingKey,               // name
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
	err = ch.Publish(
		pc.Exchange,                // exchange
		q.Name,                     // routing key
		rmqc.QueueConfig.Mandatory, // mandatory
		rmqc.QueueConfig.Immediate, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(producerJSON),
		})
	if err != nil {
		return err
	}
	log.Info("Send success.")
	return nil
}
