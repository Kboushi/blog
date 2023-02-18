# imchat
## 相关包
``` 
// 数据库
go get gorm.io/gorm
go get gorm.io/gorm/mysql

go run test/test_gorm.go

// 路由
go get -u github.com/gin-gonic/gin

// 读配置
go get github.com/spf13/viper

//swag
ginSwagger
go install github.com/swaggo/swag/cmd/swag // 1.17之后不建议用go get 下载可执行文件
swag init // 创建docs目录
swag --v //swag.exe version v1.8.10

go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

在router/app.go中引入

在service/index.go中输入
// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]

然后命令行中swag init一下
127.0.0.1/8081/swagger/index.html

在service/userService.go中输入
// GetUserList
// @Tags 首页
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]

然后命令行中swag init一下
127.0.0.1/8081/swagger/index.html

```

## 打印日志
```go
//在init sql时加入自己的logger,之后会打印在控制台
//utils/system_init.go
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
}
```





