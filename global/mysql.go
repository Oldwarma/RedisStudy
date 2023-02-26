package global

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var err error

const (
	Port     = 3306
	Name     = "redisStudy"
	UserName = "root"
	Password = "qwertyuiop"
	HOST     = "120.76.96.94"
)

func Init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Info,
		},
	)
	conn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		UserName, Password, HOST, Port, Name)
	zap.S().Info(conn)
	DB, err = gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: newLogger, //指定一个新的日志记录器，用于记录 GORM 库执行的 SQL 查询和执行时间等信息。
		NamingStrategy: schema.NamingStrategy{ //将 GORM 模型的命名转换为数据库表名
			SingularTable: true,
		},
	})
	if err != nil {
		panic("数据库链接错误:" + err.Error())
	}

	db, _ := DB.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
}
