package models

import (
	"time"

	"gorm.io/gorm"
)

type USER struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,number"`
}

type RISK_PROFILE struct {
	USER         USER `gorm:"foreignKey:UserId"`
	UserId       int
	MMPercent    float32
	BondPercent  float32
	StockPercent float32
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
