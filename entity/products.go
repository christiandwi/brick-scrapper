package entity

import (
	"github.com/christiandwi/edot/product-service/constant"
	"gorm.io/gorm"
)

type Products struct {
	ID           int64 `gorm:"column:id;primary_key"`
	ProductName  string
	Description  string
	ImageLink    string
	Price        string
	Rating       float32
	MerchantName string
	TimeScrapped string
}

func (Products) TableName() string {
	return constant.EntityProducts
}

func (u *Products) BeforeCreate(scope *gorm.DB) (err error) {
	return nil
}
