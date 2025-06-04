package model

type RateRent struct {
	Id         int     `json:"-" gorm:"column:id"`
	AmountRate int     `json:"amount_rate" gorm:"amount_rate"`
	Rate       float64 `json:"rate" gorm:"rate"`
}
