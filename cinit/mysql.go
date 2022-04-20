package cinit

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	//_ "github.com/go-sql-driver/mysql" //  mysql驱动
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"github.com/wuyadong1990/grpc-demo-user/internal/impl"
)

var Mysql *sqlx.DB
var GormDB *gorm.DB
var SqlDb *sql.DB

// 初始化连接
func mysqlInit() {
	var err error
	dataSourceName := Config.Mysql.User + ":" + Config.Mysql.Password + "@tcp(" + Config.Mysql.Addr + ":" + strconv.Itoa(Config.Mysql.Port) +
		")/" + Config.Mysql.DbName + "?parseTime=true&loc=Local"
	/*Mysql, err = sqlx.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	Mysql.SetMaxIdleConns(Config.Mysql.IDleConn)
	Mysql.SetMaxOpenConns(Config.Mysql.MaxConn)
	err = Mysql.Ping()
	if err != nil {
		panic(err)
	}*/
	GormDB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	//GormDB, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	sqlDB, err := GormDB.DB()
	// See "Important settings" section.
	sqlDB.SetConnMaxLifetime(time.Minute * 10)
	sqlDB.SetMaxOpenConns(Config.Mysql.MaxConn)
	sqlDB.SetMaxIdleConns(Config.Mysql.IDleConn)

	/*GormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: Mysql,
	}), &gorm.Config{})*/
	fmt.Printf("GormDB=%#v\n", GormDB)
	GormDB.AutoMigrate()

}

// 关闭
func mysqlClose() {
	Mysql.Close()
}
