package Zabbix

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBCONNECT *gorm.DB

func init() {
	DBCONNECT = Database("root", "Mengran@123", "zabbix", "101.35.153.11:3306")
}

func Database(username string, password string, dbname string, urlport string) (db *gorm.DB) {

	DSNCONNECT := "" + username + ":" + password + "@tcp(" + urlport + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       DSNCONNECT, // DSN data source name
		DefaultStringSize:         256,        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
		fmt.Println("数据源连接失败")
	}

	return db
}
