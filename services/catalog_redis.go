package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sun053123/go-fiber-redis/entities"
)

type catalogServiceRedis struct {
	productEnt  entities.ProductEntity
	redisClient *redis.Client
}

func NewCatalogServiceRedis(productEntity entities.ProductEntity, redisClient *redis.Client) CatalogService {
	return catalogServiceRedis{productEntity, redisClient}
}

func (serv catalogServiceRedis) GetProducts() (products []Product, err error) {

	key := "service::GetProducts"

	//get redis
	if productJson, err := serv.redisClient.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(productJson), &products) == nil {
			fmt.Println("redis")
			return products, nil
		}
	}

	// entity
	productsDB, err := serv.productEnt.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, product := range productsDB {
		products = append(products, Product{
			Name:     product.Name,
			Quantity: product.Quantity,
		})
	}
	//set redis
	if data, err := json.Marshal(products); err == nil {
		serv.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	fmt.Println("database")
	return products, nil

}
