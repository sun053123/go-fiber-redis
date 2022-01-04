package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/sun053123/go-fiber-redis/entities"
	"github.com/sun053123/go-fiber-redis/handlers"
	"github.com/sun053123/go-fiber-redis/services"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	PORT := os.Getenv("APP_PORT")

	db := initDB()
	redisClient := initRedis()
	_ = redisClient

	productEntity := entities.NewProductEntityDB(db)
	productService := services.NewCatalogServiceRedis(productEntity, redisClient)
	productHandler := handlers.NewCatalogHandler(productService)

	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		time.Sleep(time.Millisecond * 10)
		return c.SendString("Hello World")
	})

	app.Get("/products", productHandler.GetProducts)

	fmt.Printf("server ready at http://localhost%s", PORT)
	app.Listen(PORT)

}

func f() {
	PORT := os.Getenv("APP_PORT")

	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		time.Sleep(time.Millisecond * 10)
		return c.SendString("Hello World")
	})

	fmt.Printf("server ready at http://localhost%s", PORT)
	app.Listen(PORT)
}

func initDB() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", host, user, password, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
