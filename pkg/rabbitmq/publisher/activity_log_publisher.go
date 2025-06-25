package publisher

import (
	"encoding/json"
	"go-echo-clean-architecture/internal/models"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/sirupsen/logrus"

	"go-echo-clean-architecture/pkg/rabbitmq"
)

type ActivityLogPublisher struct {
	amqp     *rabbitmq.RabbitMQClient
	exchange string
}

func NewActivityLogPublisher(amqp *rabbitmq.RabbitMQClient, exchange string) *ActivityLogPublisher {

	// Menghubungkan ke RabbitMQ
	logrus.Info("*** Connecting to rabbitmq amqp ***")
	err := amqp.Connect()
	if err != nil {
		log.Fatalf("Gagal terhubung: %v", err)
	}
	logrus.Info("*** Success connected to rabbitmq amqp ***")

	err = amqp.DeclareExchange(exchange, "topic", true, false, false, false)
	if err != nil {
		log.Fatalf("Gagal mendeklarasikan exchange: %v", err)
	}

	queue, err := amqp.DeclareQueue(exchange, true, false, false, false)
	if err != nil {
		log.Fatalf("Gagal mendeklarasikan queue: %v", err)
	}

	err = amqp.BindQueue(queue.Name, exchange, exchange)
	if err != nil {
		log.Fatalf("Gagal binding queue: %v", err)
	}

	// Hanya ambil satu task pada satu waktu
	err = amqp.SetQoS(1, 0, false)
	if err != nil {
		log.Fatalf("Gagal mengatur QoS: %v", err)
	}

	return &ActivityLogPublisher{amqp: amqp, exchange: exchange}
}

func (receiver *ActivityLogPublisher) Publish(ctx echo.Context, linkId uuid.UUID, userId uuid.UUID) {
	accessLog := models.AccessLog{
		ID:         uuid.New(),
		AccessTime: time.Now(),
		IpAddress:  ctx.RealIP(),
		ClientInfo: ctx.Request().Host,
		LinkID:     linkId,
		UserID:     userId,
	}

	accessLogJson, err := json.Marshal(accessLog)
	if err != nil {
		logrus.Error("error marshal accessLogJson")
		return
	}

	logrus.Infof("publishing log activity to rabbitmq %s", string(accessLogJson))
	err = receiver.amqp.Publish("log_activity", "log_activity", true, false, accessLogJson)
	if err != nil {
		logrus.Errorf("error when publishing to rabbitmq %s", err)
	} else {
		logrus.Infof("success publishing to rabbitmq")
	}

}
