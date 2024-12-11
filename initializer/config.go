package initializer

import (
	"github.com/spf13/viper"
	"log"
)

var GlobalConfig *Config

type AppConfig struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
}

type MysqlConfig struct {
	Dsn          string `mapstructure:"dsn"`
	MaxIdleConns int    `mapstructure:"maxidleconns"`
	MaxOpenConns int    `mapstructure:"maxopenconns"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	Timeout  int    `mapstructure:"timeout"`
	Db       int    `mapstructure:"db"`
}

type Config struct {
	AppConfig   AppConfig   `mapstructure:"app"`
	MysqlConfig MysqlConfig `mapstructure:"mysql"`
	RedisConfig RedisConfig `mapstructure:"redis"`
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./initializer")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}
	//GlobalConfig 正确初始化
	GlobalConfig = &Config{}
	if err := viper.Unmarshal(GlobalConfig); err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
	}

}
