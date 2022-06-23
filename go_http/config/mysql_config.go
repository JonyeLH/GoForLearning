package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type MysqlConfig struct {
	UserName   string `yaml:"username"`
	Password   string `yaml:"password"`
	NetWork    string `yaml:"network"`
	Server     string `yaml:"server"`
	Port       string `yaml:"port"`
	DataBase   string `yaml:"database"`
	MaxConn    int    `yaml:"max_conn"`
	MaxIdeConn int    `yaml:"max_ide_conn"`
}

var DbCon *gorm.DB
var mysqlSelf MysqlConfig

func InitMysql() {
	_ = initConf(BathPath+"mysql_conf.yml", &mysqlSelf)
	fmt.Println(mysqlSelf)
	dbCon, err := NewMysql(mysqlSelf)
	if err != nil {
		panic("Mysql初始化失败！")
	}
	DbCon = dbCon
	log.Printf("mysql初始化成功！server:%s port:%s", mysqlSelf.Server, mysqlSelf.Port)
}

/*
参考链接：https://gorm.io/zh_CN/docs/connecting_to_the_database.html
*/
func NewMysql(mysqlSelf MysqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlSelf.UserName, mysqlSelf.Password,
		mysqlSelf.NetWork, mysqlSelf.Server, mysqlSelf.Port, mysqlSelf.DataBase)
	dbCon, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("mysql 初始化失败", err.Error())
		return nil, err
	}
	sql_db, err := dbCon.DB()
	if err != nil {
		log.Printf("mysql 初始化失败", err.Error())
		return nil, err
	}
	sql_db.SetMaxIdleConns(mysqlSelf.MaxConn) //最大连接数
	sql_db.SetMaxIdleConns(mysqlSelf.MaxIdeConn)
	sql_db.SetConnMaxLifetime(60) //最大生存时间(s)
	return dbCon, nil
}
