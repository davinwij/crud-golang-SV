package models

type USER struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,number"`
}

type RISK_PROFILE struct {
	MMPercent    float32
	BondPercent  float32
	StockPercent float32
	UserID       int
}

type USER_DETAIL struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Age          int     `json:"age"`
	MMPercent    float32 `json:"mmpercent"`
	BondPercent  float32 `json:"bondpercent"`
	StockPercent float32 `json:"stockpercent"`
}

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
