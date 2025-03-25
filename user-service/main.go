package main

import (
	"fmt"
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
	"github.com/Lets-Golang/Travel-Booking/user-service/repositories"
	"github.com/Lets-Golang/Travel-Booking/user-service/services"

	"github.com/golang-migrate/migrate/v4"
	migratedb "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func init() {
	godotenv.Load()
}

var log = logrus.New()

func runMigrations(db *gorm.DB, dsn string) error {
	dbInstance, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}

	driver, err := migratedb.WithInstance(dbInstance, &migratedb.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to initialize migration: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %v", err)
	}

	log.Info("Migrations applied successfully")
	return nil
}

func main() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := runMigrations(db, dsn); err != nil {
		log.Fatal("Failed to apply migrations:", err)
	}

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
