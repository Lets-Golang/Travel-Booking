package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Lets-Golang/Travel-Booking/user-service/controllers"
	"github.com/Lets-Golang/Travel-Booking/user-service/docs"
	"github.com/Lets-Golang/Travel-Booking/user-service/models"
	"github.com/Lets-Golang/Travel-Booking/user-service/repositories"
	"github.com/Lets-Golang/Travel-Booking/user-service/services"
)

func init() {
    godotenv.Load()
}

var log = logrus.New()

func main() {
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") +
        " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") +
        " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
	db.AutoMigrate(&models.UserEntity{})

	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)

	r := gin.Default()

	docs.SwaggerInfo.Title = "User Service API"
    docs.SwaggerInfo.Description = "API for managing users in the Travel Booking system"
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "localhost:8081"
    docs.SwaggerInfo.BasePath = "/"
    docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/users", controllers.CreateUser(service))
	r.GET("/users/:id", controllers.GetUser(service))

	log.Info("Starting User Service on :8081")
	r.Run(":8081")
}