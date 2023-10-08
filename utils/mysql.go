package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// 定义一个全局对象Db
	Db *sql.DB
	//定义数据库连接的相关参数值
	//连接数据库的用户名
	userName string = "root"
	//连接数据库的密码
	password string = "3953"
	//连接数据库的地址
	ipAddress string = "127.0.0.1"
	//连接数据库的端口号
	port int = 3306
	//连接数据库的具体数据库名称
	dbName string = "gotest"
	//连接数据库的编码格式
	charset string = "utf8"
)

func InitDB() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
	//Open打开一个driverName指定的数据库，dataSourceName指定数据源
	//不会校验用户名和密码是否正确，只会对dsn的格式进行检测
	Db, err = sql.Open("mysql", dsn)
	if err != nil { //dsn格式不正确的时候会报错
		return err
	}
	//尝试与数据库连接，校验dsn是否正确
	err = Db.Ping()
	if err != nil {
		fmt.Println("校验失败,err", err)
		return err
	}
	// 设置最大连接数
	Db.SetMaxOpenConns(50)
	// 设置最大的空闲连接数
	// db.SetMaxIdleConns(20)
	fmt.Println("连接数据库成功！")
	return nil
}
