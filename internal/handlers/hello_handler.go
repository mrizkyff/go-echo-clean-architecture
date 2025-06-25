package handlers

import (
	"context"
	"encoding/json"
	"go-echo-clean-architecture/internal/dto/response"
	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/services"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type HelloHandler struct {
	helloService *services.HelloService
	redisClient  *redis.Client
}

func NewHelloHandler(helloService *services.HelloService, redisClient *redis.Client) *HelloHandler {
	return &HelloHandler{
		helloService: helloService,
		redisClient:  redisClient,
	}
}

func (handler *HelloHandler) Hello(c echo.Context) error {
	message := handler.helloService.SayHello()
	ctx := context.Background()

	user := models.User{
		ID:       uuid.New(),
		Username: "Ian",
		FullName: "Homelabs",
		Password: "secret",
		Email:    "ianhomelabs@gmail.com",
		Role:     "ADMIN",
		Phone:    "+6212312332111",
		Address:  "Indonesia",
		Age:      20,
		Links:    nil,
		Auditable: models.Auditable{
			CreatedAt: time.Now(),
			CreatedBy: uuid.New(),
			UpdatedAt: time.Now(),
			UpdatedBy: uuid.New(),
			DeletedAt: nil,
			DeletedBy: uuid.New(),
		},
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		logrus.Error(err)
	}

	result, err := handler.redisClient.Set(ctx, "sample:1", userJson, 0).Result()
	if err != nil {
		logrus.Error("redis set err: ", err)
	}
	logrus.Info(result)

	return response.Success(c, 200, message, nil)
}
