package config

import (
	"fmt"
	"gorm.io/gorm"
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

var Dbcon *gorm.DB
var mysqlSelf MysqlConfig

func InitMysql() {
	_ = initConf(BathPath+"mysql_conf.yaml", mysqlSelf)
	fmt.Println(mysqlSelf)
}
