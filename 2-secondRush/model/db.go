package model

import "RedisStudy/global"

func InitDB() {
	err := global.DB.AutoMigrate(&Voucher{}, &VoucherOrder{})
	if err != nil {
		panic("创建Voucher表失败:" + err.Error())
	}
}
