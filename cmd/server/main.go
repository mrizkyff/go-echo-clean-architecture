package main

import (
	"go-echo-clean-architecture/internal/handlers"
	"go-echo-clean-architecture/internal/middlewares"
	"go-echo-clean-architecture/internal/repositories"
	"go-echo-clean-architecture/internal/routes"
	"go-echo-clean-architecture/internal/services"
	"go-echo-clean-architecture/internal/utils"
	"go-echo-clean-architecture/pkg/config"
	"go-echo-clean-architecture/pkg/database"
	"go-echo-clean-architecture/pkg/rabbitmq"
	"go-echo-clean-architecture/pkg/rabbitmq/consumer"
	"go-echo-clean-architecture/pkg/rabbitmq/publisher"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"go.elastic.co/apm/module/apmechov4/v2"
	"gorm.io/gorm"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	// Initialize echo
	e := echo.New()

	// Apm
	e.Use(apmechov4.Middleware())

	// Initialize gorm
	var gormDb *gorm.DB

	// Initialize config
	// Read Config
	remoteConfig := config.NewRemoteConfig()
	appConfig := config.NewAppConfig(remoteConfig)

	// Initialize client
	jwtUtil := utils.NewJWTUtil(appConfig.JwtConfig)
	amqpClient := rabbitmq.NewRabbitMQClient(appConfig.RabbitMqConfig)
	postgresClient := database.NewPostgresClient(gormDb, appConfig.PostgresConfig)
	redisClient := database.NewRedisClient(appConfig.RedisConfig)

	// Initialize rabbitmq publisher
	logActivityPublisher := publisher.NewActivityLogPublisher(amqpClient, "log_activity")

	//Initialize auth middleware
	authMiddleware := middlewares.NewAuthMiddleware(jwtUtil)

	// Register middleware
	e.Use(middlewares.LoggerMiddleware()) // Tambahkan middleware logger
	e.Use(middlewares.RecoverMiddleware())
	e.Use(middlewares.ConfigureCORS())
	e.Use(middlewares.ErrorHandlingMiddleware())
	e.Use(middlewares.NotFoundHandler())
	//e.Use(authMiddleware.Authenticate())

	// Inisialisasi koneksi database
	if err := postgresClient.InitDBConnection(); err != nil {
		logrus.Fatalf("Gagal menginisialisasi database: %v", err)
	}

	// Get database connection with SQL logging enabled
	db := postgresClient.GetDB()

	// Log database connection info
	logrus.Info("Database connection established with SQL logging enabled")

	// Initialize repositories
	userRepository := repositories.NewUserRepositoryImpl(db)
	linkRepository := repositories.NewLinkRepositoryImpl(db)
	accessLogRepository := repositories.NewAccessLogRepositoryImpl(db)
	userRedisRepository := repositories.NewUserRedisRepositoryImpl(&redisClient)

	// Initialize services
	helloService := services.NewHelloService()
	userService := services.NewUserServiceImpl(userRepository)
	linkService := services.NewLinkServiceImpl(linkRepository, userRepository)
	accessLogService := services.NewAccessLogService(accessLogRepository)
	authenticationService := services.NewAuthenticationServiceImpl(userRepository, userService, jwtUtil, userRedisRepository)

	// Initialize handlers
	helloHandler := handlers.NewHelloHandler(helloService, &redisClient)
	userHandler := handlers.NewUserHandler(userService, jwtUtil)
	linkHandler := handlers.NewLinkHandler(linkService, logActivityPublisher, jwtUtil)
	accessLogHandler := handlers.NewAccessLogHandler(*accessLogService)
	authenticationHandler := handlers.NewAuthenticationHandler(authenticationService, userService)

	// Register routes
	routes.RegisterHelloRoutes(helloHandler, e)
	routes.RegisterUserRoutes(userHandler, e, authMiddleware)
	routes.RegisterLinkRoutes(linkHandler, e, authMiddleware)
	routes.RegisterAccessLogRoutes(accessLogHandler, e)
	routes.RegisterAuthenticationRoute(authenticationHandler, e)

	// Register catch-all route for 404s (should be after all other routes)
	middlewares.RegisterNotFoundRoute(e)

	e.GET("/", func(c echo.Context) error {
		return c.File("static/index.html")
	})

	// Initialize rabbitmq consumer
	logActivityConsumer := consumer.NewActivityLogConsumer(amqpClient, "log_activity", accessLogService)
	// Jalankan consumer dalam goroutine
	logrus.Infof("Starting to consume messages from activity_log queue...")
	go logActivityConsumer.Consume()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))

	// Defer function
	defer func() {
		err := postgresClient.CloseDB()
		if err != nil {

		}
	}()
}
