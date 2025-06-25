package rabbitmq

import (
	"context"
	"fmt"
	"go-echo-clean-architecture/internal/models/config"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQClient adalah struct untuk mengelola koneksi RabbitMQ
type RabbitMQClient struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	uri          string
	exchangeName string
	exchangeType string
	queueName    string
}

// NewRabbitMQClient membuat instance baru RabbitMQClient
func NewRabbitMQClient(config config.RabbitMQConfig) *RabbitMQClient {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", config.Username, config.Password, config.Host, config.Port, config.Vhost)
	return &RabbitMQClient{
		uri: uri,
	}
}

// Connect membuat koneksi ke server RabbitMQ
func (r *RabbitMQClient) Connect() error {
	var err error
	r.conn, err = amqp.Dial(r.uri)
	if err != nil {
		return fmt.Errorf("gagal terhubung ke RabbitMQ: %w", err)
	}

	r.channel, err = r.conn.Channel()
	if err != nil {
		return fmt.Errorf("gagal membuka channel: %w", err)
	}

	return nil
}

// Close menutup koneksi dan channel RabbitMQ
func (r *RabbitMQClient) Close() {
	if r.channel != nil {
		err := r.channel.Close()
		if err != nil {
			return
		}
	}
	if r.conn != nil {
		err := r.conn.Close()
		if err != nil {
			return
		}
	}
}

// DeclareExchange mendeklarasikan exchange
func (r *RabbitMQClient) DeclareExchange(name, exchangeType string, durable, autoDelete, internal, noWait bool) error {
	r.exchangeName = name
	r.exchangeType = exchangeType
	return r.channel.ExchangeDeclare(
		name,         // name
		exchangeType, // type
		durable,      // durable
		autoDelete,   // auto-deleted
		internal,     // internal
		noWait,       // no-wait
		nil,          // arguments
	)
}

// DeclareQueue mendeklarasikan queue
func (r *RabbitMQClient) DeclareQueue(name string, durable, autoDelete, exclusive, noWait bool) (amqp.Queue, error) {
	r.queueName = name
	return r.channel.QueueDeclare(
		name,       // name
		durable,    // durable
		autoDelete, // delete when unused
		exclusive,  // exclusive
		noWait,     // no-wait
		nil,        // arguments
	)
}

// BindQueue melakukan binding queue ke exchange
func (r *RabbitMQClient) BindQueue(queueName, routingKey, exchangeName string) error {
	return r.channel.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
}

// Publish mengirim pesan ke exchange
func (r *RabbitMQClient) Publish(exchange, routingKey string, mandatory, immediate bool, msg []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.channel.PublishWithContext(
		ctx,
		exchange,   // exchange
		routingKey, // routing key
		mandatory,  // mandatory
		immediate,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
}

// Consume menerima pesan dari queue
func (r *RabbitMQClient) Consume(queueName, consumerName string, autoAck, exclusive, noLocal, noWait bool) (<-chan amqp.Delivery, error) {
	return r.channel.Consume(
		queueName,    // queue
		consumerName, // consumer
		autoAck,      // auto-ack
		exclusive,    // exclusive
		noLocal,      // no-local
		noWait,       // no-wait
		nil,          // args
	)
}

// SetQoS mengatur prefetch count
func (r *RabbitMQClient) SetQoS(prefetchCount, prefetchSize int, global bool) error {
	return r.channel.Qos(
		prefetchCount, // prefetch count
		prefetchSize,  // prefetch size
		global,        // global
	)
}
