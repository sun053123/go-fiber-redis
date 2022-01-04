package entities

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

//ไม่ต้องการให้ struct ที่รับข้อมูลจาก database expose ออกไป เหมือนการทำ access modifier (java, c#)
type product struct {
	ID       int
	Name     string
	Quantity int
}

type ProductEntity interface {
	GetProducts() ([]product, error)
}

func mockData(db *gorm.DB) error {
	var count int64
	db.Model(&product{}).Count(&count)
	if count > 0 {
		return nil
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	products := []product{}
	for i := 0; i < 5000; i++ {
		products = append(products, product{
			Name:     fmt.Sprintf("Product%v", i+1),
			Quantity: random.Intn(100),
		})
	}
	return db.Create(&products).Error
}
