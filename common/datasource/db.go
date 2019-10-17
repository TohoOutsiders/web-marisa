/**
 * @Author: Tomonori
 * @Date: 2019/6/18 14:56
 * @File: db
 * @Desc:
 */
package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"server/common/setting"
)

type Db struct {
	Conn *gorm.DB
}

func (d *Db) Connect() error {
	conf := setting.Config.Database
	var (
		dbType = conf.Type
		dbName = conf.Name
		user   = conf.User
		pwd    = conf.Password
		host   = conf.Host
	)

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, dbName))
	if err != nil {
		log.Fatal("connecting mysql error: ", err)
		return err
	}

	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	d.Conn = db

	log.Println("Connect Mysql Success")

	return nil
}

func (d *Db) DB() *gorm.DB {
	return d.Conn
}
