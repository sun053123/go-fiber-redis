package entities

import (
	"gorm.io/gorm"
)

type productEntityDB struct {
	db *gorm.DB
}

//เปรียบเสมือน constructure ให้คนที่จะมาใช้ได้ต้องส่ง db เข้ามาก่อน
func NewProductEntityDB(db *gorm.DB) ProductEntity {
	db.AutoMigrate(&product{})
	mockData(db)
	return productEntityDB{db: db}
}

// จะทำตัวเป็น method ของ struct productEntityDB จึงใช้ reciver func เปรียบเสมือน method ใน class นั้น
func (ent productEntityDB) GetProducts() (products []product, err error) {
	err = ent.db.Order("quantity desc").Limit(30).Find(&products).Error
	return products, err
}
