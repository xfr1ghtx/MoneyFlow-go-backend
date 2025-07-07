package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stepanpotapov/moneyflow-go-backend/internal/handler"
	"github.com/stepanpotapov/moneyflow-go-backend/internal/repository"
	"github.com/stepanpotapov/moneyflow-go-backend/internal/service"

	// Swagger
	_ "github.com/stepanpotapov/moneyflow-go-backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Загрузка .env
	"github.com/joho/godotenv"
)

// @title MoneyFlow API
// @version 1.0
// @description API для регистрации, логина и логаута пользователей
// @BasePath /
// @schemes http
// @host localhost:8080
func main() {
	// Загружаем переменные окружения из .env, если файл существует
	_ = godotenv.Load()

	// Получаем строку подключения к базе данных из переменной окружения или используем значение по умолчанию
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		dbUrl = "postgres://moneyflow_user:moneyflow_pass@localhost:5432/moneyflow?sslmode=disable"
	}

	// Устанавливаем соединение с базой данных с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer pool.Close()

	// Инициализируем репозитории, сервис и обработчик для авторизации
	repo := repository.NewUserRepository(pool)
	refreshRepo := repository.NewRefreshTokenRepository(pool)
	authService := service.NewAuthService(repo, refreshRepo, os.Getenv("JWT_SECRET"))
	authHandler := handler.NewAuthHandler(authService)

	// Создаём новый роутер Gin с логированием и обработкой паник
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Регистрируем маршруты для регистрации, логина и логаута
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/logout", authHandler.Logout)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health-check endpoint (не документируется в Swagger)
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Простой healthcheck endpoint (оставлен для обратной совместимости)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Запускаем HTTP сервер на порту 8080
	r.Run(":8080")
}
