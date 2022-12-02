package main

import (
	"context"
	"fmt"
	"time"

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
	clientConnectionName := "test123"
	rabbitMQ := New(clientConnectionName)
	consumerTag := "Tag123"
	deliveries, err := rabbitMQ.NewConsume(rabbitMQSetting.TestQueueName, consumerTag)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(fmt.Sprintf("rabbit tag %s consumer start success", consumerTag))

	go func() {
		for d := range deliveries {
			fmt.Println("start delivery")
			time.Sleep(20 * time.Second)
			fmt.Println(fmt.Sprintf(
				"got %dB delivery: [%v] %q",
				len(d.Body),
				d.DeliveryTag,
				d.Body,
			))
			d.Ack(false)
		}
	}()
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

func (service *RabbitMQ) NewConsume(queueName, consumerTag string) (<-chan amqp.Delivery, error) {
	channel := service.Channel
	deliveries, err := channel.Consume(
		queueName,   // name
		consumerTag, // consumerTag,
		false,       // autoAck
		false,       // exclusive
		false,       // noLocal
		false,       // noWait
		nil,         // arguments
	)
	return deliveries, err
}

func New(clientConnectionName string) RabbitMQ {
	rabbitMQSetting := RabbitMQSetting{
		Uri:          "",
		ExchangeName: "taikang",
		ExchangeType: "topic",
	}
	uri := rabbitMQSetting.Uri
	config := amqp.Config{Properties: amqp.NewConnectionProperties()}
	config.Properties.SetClientConnectionName(clientConnectionName)
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
