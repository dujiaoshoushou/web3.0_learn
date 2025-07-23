package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var logWriter io.Writer

func init() {
	InitDb()
	//InitAutoMigrate()
}

func InitAutoMigrate() {
	if err := DB.AutoMigrate(&Comment{}, &Post{}, &User{}); err != nil {
		log.Fatalln(err)
	}
}

func InitDb() {
	var connectionString string = "root:123456@tcp(127.0.0.1:3307)/gorm_example?charset=utf8mb4&parseTime=True&loc=Local"
	logWriter, _ = os.OpenFile("./sql.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	customerLogger := logger.New(log.New(logWriter, "\r\n", log.LstdFlags), logger.Config{
		LogLevel:                  logger.Info,
		Colorful:                  false,
		IgnoreRecordNotFoundError: false,
		ParameterizedQueries:      false,
		SlowThreshold:             200 * time.Millisecond,
	})

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: customerLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_unicorn",
			SingularTable: true,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		log.Fatal("初始化数据库连接异常")
	}
	DB = db
}
