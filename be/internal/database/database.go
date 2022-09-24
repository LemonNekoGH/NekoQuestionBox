package database

import (
	"errors"
	"fmt"
	"neko-question-box-be/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "embed"
)

var DB *gorm.DB

//go:embed questions.sql
var questionSql string

// 初始化数据库
func InitDB() {
	dialector := postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		config.Conf.Database.Host,
		config.Conf.Database.Username,
		config.Conf.Database.Password,
		config.Conf.Database.Database,
		config.Conf.Database.Port,
		config.Conf.Database.SSLMode,
	))
	var err error
	DB, err = gorm.Open(dialector)
	if err != nil {
		panic(err)
	}
}

func InitDBTest() {
	InitDB()
	// 删除然后重新建表
	err := DB.Exec(questionSql).Error
	if err != nil {
		panic(err)
	}
}

// 是否是「记录未找到」错误
func IsNoRecordFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
