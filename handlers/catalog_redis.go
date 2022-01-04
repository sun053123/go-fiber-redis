package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/sun053123/go-fiber-redis/services"
)

type catalogHandlerRedis struct {
	catalogSrv  services.CatalogService
	redisClient *redis.Client
}

func NewCatalogHandlerRedis(catalogSrv services.CatalogService, redisClient *redis.Client) CatalogHandler {
	return catalogHandlerRedis{catalogSrv, redisClient}
}

func (handl catalogHandlerRedis) GetProducts(c *fiber.Ctx) error {

	key := "handler::GetProducts"

	//redis get
	if reponseJson, err := handl.redisClient.Get(context.Background(), key).Result(); err == nil {
		fmt.Println("redis")
		c.Set("Content-Type", "application/json")
		return c.SendString(reponseJson)
	}

	//service
	products, err := handl.catalogSrv.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":  "ok",
		"product": products,
	}

	//redis set
	if data, err := json.Marshal(response); err == nil {
		handl.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	fmt.Println("database")
	return c.JSON(response)
}
