package main

import (
	"golang-interface/demo/mq"
	"log"
)

type MyConsumer struct{}

func (mc MyConsumer) Receive(msg string) {
	log.Println(msg)
	return
}

func main() {
	rc1 := mq.ReceiveConf{QueueValue: "QV-1111111111", Exchange: "Exchange.1111111.1", RoutingKey: "RoutingKey1"}
	go rc1.Consumer(&MyConsumer{})
	rc2 := mq.ReceiveConf{QueueValue: "QV-2222222222", Exchange: "Exchange.2222222.2", RoutingKey: "RoutingKey2"}
	go rc2.Consumer(&MyConsumer{})
	var c = make(chan bool)
	<-c
	log.Println("... End ...")
}
