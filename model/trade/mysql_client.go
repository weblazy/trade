package trade

import (
	"fmt"

	"github.com/weblazy/easy/db/emysql"
	"gorm.io/gorm"
)

const TradeMysql = "Trade"

func GetDB() *gorm.DB {
	return emysql.GetMysql(TradeMysql).DB
}

func SchemaMigrate() {
	fmt.Println("开始初始化trade数据库")
	//自动建表，数据迁移
	_ = GetDB().Set("gorm:table_options", "CHARSET=utf8mb4 comment='用户表' AUTO_INCREMENT=1;").AutoMigrate(&User{})

	fmt.Println("数据库trade初始化完成")
}
