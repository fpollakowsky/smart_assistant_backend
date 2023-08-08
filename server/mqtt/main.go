package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"home-ai-backend/server/config"
	"log"
	"math/rand"
	"os"
	"time"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func Connect() mqtt.Client {
	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	//mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)

	rand.Seed(time.Now().UnixNano())

	opts := mqtt.NewClientOptions().AddBroker("tcp://" + config.BROKER_IP + ":" + config.BROKER_PORT).SetClientID("nethcon" + randSeq(8))

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// subscribe(c)
	return c
}

func NewRequest(channel, room, payload string) {
	var client = Connect()
	client.Publish("/nethcon/"+room+"/"+channel, 1, false, payload)
	client.Disconnect(250)
}

func subscribe(client mqtt.Client, topic string) {
	if token := client.Subscribe(topic, 0, messagePubHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
