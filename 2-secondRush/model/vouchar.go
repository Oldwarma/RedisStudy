package model

import (
	"RedisStudy/global"
	"time"
)

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
	result := global.DB.First(&v)
	if result.Error != nil {
		return v, result.Error
	} else {
		return v, nil
	}
}

func (r Voucher) Add(amount int, startTime, endTime time.Time) (int, error) {
	voucher := Voucher{
		Amount:     amount,
		StartTime:  startTime,
		EndTime:    endTime,
		IsValid:    1,
		CreateTime: time.Now(),
	}
	res := global.DB.Save(&voucher)
	if res.Error != nil {
		return 0, res.Error
	} else {
		return voucher.ID, nil
	}
}

func (r Voucher) DecreaseStock(id int) (int, error) {
	res, err := r.GetById(id)
	if err != nil {
		return 0, err
	}
	global.DB.Model(Voucher{}).Where("id=?", id).Update("amount", int32(res.Amount-1))
	return id, nil
}
