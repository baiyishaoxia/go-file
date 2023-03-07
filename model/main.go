package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-file/common"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var DB *gorm.DB

func createAdminAccount() {
	var user User
	DB.Where(User{Role: common.RoleAdminUser}).Attrs(User{
		Username:    "admin",
		Password:    "123456",
		Role:        common.RoleAdminUser,
		Status:      common.UserStatusEnabled,
		DisplayName: "Administrator",
	}).FirstOrCreate(&user)
}

func CountTable(tableName string) (num int64) {
	DB.Table(tableName).Count(&num)
	return
}

func InitDB(configInfo *common.ConfigModel) (db *gorm.DB, err error) {
	mysqlInfo := configInfo.MySql
	if mysqlInfo != nil {
		// Use MySQL
		host := mysqlInfo.Host
		port := mysqlInfo.Port
		user := mysqlInfo.User
		password := mysqlInfo.Password
		dbName := mysqlInfo.Dbname
		prefix := mysqlInfo.Prefix
		maxIdleConn := mysqlInfo.MaxIdleConn
		maxOpenConn := mysqlInfo.MaxOpenConn
		dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, dbName)
		db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   prefix, // 表前缀
				SingularTable: true,   // 禁用表名复数
			}})
		sqlDb, _ := db.DB()
		sqlDb.SetMaxIdleConns(maxIdleConn)
		sqlDb.SetMaxOpenConns(maxOpenConn)
	} else {
		// Use SQLite
		db, err = gorm.Open(sqlite.Open(configInfo.Server.SqlitePath), &gorm.Config{})
	}
	if err == nil {
		DB = db
		_ = db.AutoMigrate(&File{})
		_ = db.AutoMigrate(&Image{})
		_ = db.AutoMigrate(&User{})
		_ = db.AutoMigrate(&Option{})
		createAdminAccount()
		return DB, err
	} else {
		log.Fatal(err)
	}
	return nil, err
}
