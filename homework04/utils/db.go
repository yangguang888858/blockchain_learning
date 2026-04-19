package utils

import (
	"fmt"
	"homework04/config"
	"homework04/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	fmt.Println("执行db.go中的init()函数...")
	//连接数据库
	db, err := Connect()
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}
	//创建表
	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		fmt.Println("创建表失败", err)
		return
	}
}

func Connect() (*gorm.DB, error) {

	//连接myql数据库
	config, err := config.ReadConfig()
	if err != nil {
		fmt.Println("读取数据库配置失败", err)
		return nil, err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.UserName,
		config.Database.PassWord,
		config.Database.Host,
		config.Database.Port,
		config.Database.DbName,
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return nil, err
	}
	//打印sql
	db.Logger = db.Logger.LogMode(logger.Info)
	return db, nil
}
