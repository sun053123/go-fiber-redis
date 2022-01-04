package entities

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productEntityRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductEntityRedis(db *gorm.DB, redisClient *redis.Client) ProductEntity {
	db.AutoMigrate(&product{})
	mockData(db)
	return productEntityRedis{db, redisClient}
}

func (ent productEntityRedis) GetProducts() (products []product, err error) {

	key := "entity::GetProducts"
	//redis get
	productsJson, err := ent.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productsJson), &products)
		if err == nil {
			fmt.Println("redis")
			return products, nil
		}
	}

	//database
	err = ent.db.Order("quantity desc").Limit(30).Find(&products).Error
	if err != nil {
		return nil, err
	}

	//redis set
	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}
	ent.redisClient.Set(context.Background(), key, string(data), time.Second*10).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("database")
	return products, nil
}
