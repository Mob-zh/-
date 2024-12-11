package initializer

import (
	"attendance_uniapp/global"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func initMysql() {
	// TODO: 初始化数据库连接

	dsn := GlobalConfig.MysqlConfig.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql db: %v", err)
	}
	sqlDB.SetMaxIdleConns(GlobalConfig.MysqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(GlobalConfig.MysqlConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	global.DB = db
}

func initRedis() {

	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:        GlobalConfig.RedisConfig.Addr,                                 // Redis 服务器地址
		Password:    GlobalConfig.RedisConfig.Password,                             // Redis 密码
		DB:          GlobalConfig.RedisConfig.Db,                                   // Redis 数据库
		DialTimeout: time.Duration(GlobalConfig.RedisConfig.Timeout) * time.Second, // 连接超时时间
	})

	// 测试 Redis 连接
	_, err := global.RedisClient.Ping(global.RedisCtx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")
}
