package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")    // yaml的名字
	viper.AddConfigPath("config") // 配置的位置
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
	log.Println("config app:", viper.Get("app"))
	log.Println("config mysql:", viper.Get("mysql"))
}

func InitMySQL() {
	// 自定义日志模板，打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        // 彩色
		},
	)

	var err error
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Mysql inited ....")
	// user := models.UserBasic{}
	// DB.Find(&user)
	// log.Println(user)
}
