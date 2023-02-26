package model

import "time"

//秒杀

type Voucher struct {
	ID         int       `json:"id" gorm:"primary_key;type:int"`
	Amount     int       `json:"amount" gorm:"amount"`
	StartTime  time.Time `json:"startTime" gorm:"start_time"`
	EndTime    time.Time `json:"endTime" gorm:"end_time"`
	CreateTime time.Time `gorm:"create_time;default:null"`
	UpdateTime time.Time `gorm:"update_time;default;null"`
	IsValid    int       `gorm:"is_valid"`
}

func (r Voucher) GetById(id int) (Voucher, error) {
	var v = Voucher{
		ID: id,
	}

}
