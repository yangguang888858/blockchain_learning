package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstruct:"server"`
	Database DatabaseConfig `mapstruct:"database"`
	Jwt      JwtConfig      `mapstruct:"jwt"`
}

type ServerConfig struct {
	Host string `mapstruct:"host"`
	Port string `mapstruct:"port"`
	Mode string `mapstruct:"mode"`
}

type DatabaseConfig struct {
	Host     string `mapstruct:"host"`
	Port     string `mapstruct:"port"`
	UserName string `mapstruct:"userName"`
	PassWord string `mapstruct:"passWord"`
	DbName   string `mapstruct:"dbName"`
}

type JwtConfig struct {
	Secret string `mapstruct:"secret"`
	Expire string `mapstruct:"expire"`
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") //注意配置路径 .是当前路径

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取配置失败", err)
		return
	}
}

func ReadConfig() (*Config, error) {
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("解析配置失败", err)
		return nil, err
	}
	return &config, nil
}
