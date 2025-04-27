package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (

	// MysqlInterface is an interface that represent mysql methods in package database
	MysqlInterface interface {
		OpenMysqlConn() (*MySQL, error)
	}

	// MySQL return gorm mysql client
	MySQL struct {
		Client *gorm.DB
	}

	// MysqlConfig is a struct to map given config
	MysqlConfig struct {
		DatabaseHost     string
		DatabaseUser     string
		DatabasePassword string
		DatabasePort     string
		DatabaseName     string
	}
)

func NewMysql(c MysqlConfig) MysqlInterface {
	return &MysqlConfig{
		DatabaseHost:     c.DatabaseHost,
		DatabaseUser:     c.DatabaseUser,
		DatabasePassword: c.DatabasePassword,
		DatabasePort:     c.DatabasePort,
		DatabaseName:     c.DatabaseName,
	}
}

func (c MysqlConfig) OpenMysqlConn() (rds *MySQL, err error) {

	fmt.Printf("Start open mysql connection...")

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseName,
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &MySQL{Client: db}, nil
}
