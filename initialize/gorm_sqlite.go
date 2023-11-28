package initialize

import (
	"github.com/FateBug403/FateBugGUI/global"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)
// InitGormSqlite 初始化数据库
func InitGormSqlite() error{
	if db, err := gorm.Open(sqlite.Open("example.db"), &gorm.Config{}); err != nil {
		return err
	} else {
		global.GVA_DB = db
		if global.GVA_DB != nil {
			err = RegisterTables()
			if err != nil {
				return err
			}
		}
		return nil
	}
}
// RegisterTables 注册数据库表
func RegisterTables() error {
	var err error
	db := global.GVA_DB
	err = db.AutoMigrate(
	)
	if err != nil {
		return err
	}
	return nil
}