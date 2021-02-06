package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func Connect() mqtt.Client {
	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	//mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)

	var broker = "192.168.2.114"
	var port = "1883"

	opts := mqtt.NewClientOptions().AddBroker("tcp://" + broker + ":" + port).SetClientID("sample")

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// subscribe(c)
	return c
}

func NewRequest(client mqtt.Client, channel, room, payload string) {
	client.Publish("/nethcon/"+room+"/"+channel, 1, false, payload)
}

func subscribe(client mqtt.Client, topic string) {
	if token := client.Subscribe(topic, 0, messagePubHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}
