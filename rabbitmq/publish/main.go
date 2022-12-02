package main

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	Setup()

	rabbitMQSetting := RabbitMQSetting{
		Uri:                 "",
		ExchangeName:        "taikang",
		ExchangeType:        "topic",
		TestQueueName:       "testQueueName",
		TestQueueBindingKey: "testQueueBindingKey",
	}

	rabbitMQ := New()
	for i := 0; i <= 10; i++ {
		rabbitMQ.PublishWithContext(context.Background(), rabbitMQSetting.ExchangeName,
			rabbitMQSetting.TestQueueBindingKey, fmt.Sprintf("%d test", i))
	}
	select {}

}

type RabbitMQSetting struct {
	Uri                 string `json:"uri"`
	ExchangeName        string `json:"exchangeName"`
	ExchangeType        string `json:"exchangeType"`
	TestQueueName       string `json:"testQueueName"`
	TestQueueBindingKey string `json:"testQueueBindingKey"`
}

func Setup() {
	rabbitMQSetting := RabbitMQSetting{
		Uri:                 "",
		ExchangeName:        "taikang",
		ExchangeType:        "topic",
		TestQueueName:       "testQueueName",
		TestQueueBindingKey: "testQueueBindingKey",
	}

	exchangeType := rabbitMQSetting.ExchangeType
	exchangeName := rabbitMQSetting.ExchangeName
	uri := rabbitMQSetting.Uri
	testQueueName := rabbitMQSetting.TestQueueName
	testQueueBindingKey := rabbitMQSetting.TestQueueBindingKey

	config := amqp.Config{Properties: amqp.NewConnectionProperties()}
	config.Properties.SetClientConnectionName("sample-producer")
	connection, err := amqp.DialConfig(uri, config)

	if err != nil {
		panic(fmt.Sprintf("rabbit Dial: %s", err.Error()))
	}
	channel, err := connection.Channel()
	if err != nil {
		panic(fmt.Sprintf("Channel: %s", err.Error()))
	}

	err = channel.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)

	if err != nil {
		panic(fmt.Sprintf("Exchange Declare: %s", err))
	}

	_, err = channel.QueueDeclare(
		testQueueName, // name of the queue
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // noWait
		nil,           // arguments
	)
	if err != nil {
		panic(fmt.Sprintf("Queue Declare: %s", err))
	}
	err = channel.QueueBind(
		testQueueName,       // name of the queue
		testQueueBindingKey, // bindingKey
		exchangeName,        // sourceExchange
		false,               // noWait
		nil,                 // arguments
	)
	if err != nil {
		panic(fmt.Sprintf("Queue Bind: %s", err))
	}
	connection.Close()
}

type RabbitMQ struct {
	Channel *amqp.Channel
}

func (service *RabbitMQ) PublishWithContext(ctx context.Context, exchange, key, body string) error {
	channel := service.Channel

	err := channel.PublishWithContext(ctx,
		exchange, // publish to an exchange
		key,      // routing to 0 or more queues
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Persistent, // 1=non-persistent, 2=persistent
			Priority:        0,               // 0-9
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func New() RabbitMQ {
	rabbitMQSetting := RabbitMQSetting{
		Uri:          "",
		ExchangeName: "taikang",
		ExchangeType: "topic",
	}
	uri := rabbitMQSetting.Uri
	config := amqp.Config{Properties: amqp.NewConnectionProperties()}
	config.Properties.SetClientConnectionName("sample-producer")
	connection, err := amqp.DialConfig(uri, config)
	if err != nil {
		fmt.Println(fmt.Sprintf("rabbit Dial: %s", err.Error()))
	}
	channel, err := connection.Channel()
	if err != nil {
		fmt.Println(fmt.Sprintf("Channel: %s", err.Error()))
	}

	return RabbitMQ{
		Channel: channel,
	}
}
