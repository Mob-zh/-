package initializer

import (
	"github.com/spf13/viper"
	"log"
)

var GlobalConfig *Config

// AppConfig 定义了应用程序配置
type AppConfig struct {
	Name string
	Port string
}

// DatabaseConfig 定义了数据库配置
type DatabaseConfig struct {
	Dsn          string
	MaxIdleConns int
	MaxOpenConns int
}

// Config 组合了AppConfig和DatabaseConfig
type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

// InitConfig 加载并初始化配置
func InitConfig() {
	//TODO: 初始化配置

	viper.SetConfigName("config")        // 设置配置文件的名称和路径
	viper.AddConfigPath("./initializer") // 配置文件在根目录的config目录下
	viper.SetConfigType("yaml")          // 配置文件格式是YAML

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}

	// 将配置文件映射到结构体
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
	}

}
