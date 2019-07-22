package mq

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type ReceiveConf struct {
	QueueValue string // @Queue(value='') 相当于  kafka group

	Exchange string

	RoutingKey string // Key 相当于 kafka topic
}

func (rc ReceiveConf) String() string {
	return fmt.Sprintf("{ QueueValue : %s, Exchange : %s, RoutingKey : %s }", rc.QueueValue, rc.Exchange, rc.RoutingKey)
}

type ConsumerInterface interface {
	Receive(msg string)
}

func (rc ReceiveConf) Consumer(ci ConsumerInterface) error {
	log.Println(rc)
	i := 1
	body := `{"id":"%s","status":"SSSSSS","ownerId":"%S","orderId":"","groupType":"GROUP_BUY","productId":"P_100%d","productName":"makeup This is a very name"}`
	for {
		ci.Receive(fmt.Sprintf(body, rc.QueueValue, rc.Exchange, i))
		i++
		min := 1
		max := 10
		waitTime := rand.Intn(max-min) + min
		log.Println(waitTime)
		time.Sleep(time.Duration(waitTime) * time.Second)

	}
}
