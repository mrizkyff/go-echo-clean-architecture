package consumer

import (
	"encoding/json"
	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/services"
	"go-echo-clean-architecture/pkg/rabbitmq"

	"github.com/sirupsen/logrus"
)

type ActivityLogConsumer struct {
	amqp             *rabbitmq.RabbitMQClient
	queueName        string
	accessLogService *services.AccessLogService
}

func NewActivityLogConsumer(amqp *rabbitmq.RabbitMQClient, queueName string, accessLogService *services.AccessLogService) *ActivityLogConsumer {
	return &ActivityLogConsumer{amqp: amqp, queueName: queueName, accessLogService: accessLogService}
}

func (receiver *ActivityLogConsumer) Consume() {
	messages, err := receiver.amqp.Consume(receiver.queueName, "activity_log_consumer_golang", true, false, false, false)
	if err != nil {
		logrus.Errorf("error consuming messages %s", err)
		return
	}

	go func() {
		for msg := range messages {
			// convert msg.Body to AccessLog
			accessLog := models.AccessLog{}
			err := json.Unmarshal(msg.Body, &accessLog)
			if err != nil {
				logrus.Errorf("error unmarshal accessLog %s", err)
				continue
			}

			// Save accessLog to database
			_, err = receiver.accessLogService.Create(&accessLog)

			// Process the accessLog as needed
			logrus.Infof("Processed access log: %+v", accessLog)
		}
	}()
}
