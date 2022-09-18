package database

import (
	"errors"

	"gorm.io/gorm"
)

// 初始化数据库
func InitDB() {

}

// 是否是「记录未找到」错误
func IsNoRecordFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
