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
	"server/setting"
)

type Db struct {
	Conn *gorm.DB
}

func (d *Db) Connect() error {
	var (
		dbType, dbName, user, pwd, host string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalln("Fail to get config section 'database': ", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	pwd = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()

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

//type MysqlConnectionPool struct {
//}
//
//var instance *MysqlConnectionPool
//var once sync.Once
//
//var db *gorm.DB
//var errDB error
//
//func GetInstace() *MysqlConnectionPool {
//	once.Do(func() {
//		instance = &MysqlConnectionPool{}
//	})
//	return instance
//}
//
//func (m *MysqlConnectionPool) InitDataPool() (issue bool) {
//	var (
//		dbType, dbName, user, pwd, host string
//	)
//
//	sec, err := setting.Cfg.GetSection("database")
//	if err != nil {
//		log.Fatalln("Fail to get config section 'database': ", err)
//	}
//	dbType = sec.Key("TYPE").String()
//	dbName = sec.Key("NAME").String()
//	user = sec.Key("USER").String()
//	pwd = sec.Key("PASSWORD").String()
//	host = sec.Key("HOST").String()
//
//	db, errDB = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, dbName))
//	if errDB != nil {
//		log.Fatal("connecting mysql error: ", errDB)
//		return false
//	}
//
//	db.SingularTable(true)
//	db.DB().SetMaxIdleConns(10)
//	db.DB().SetMaxOpenConns(100)
//
//	return true
//}
//
//func (m *MysqlConnectionPool) GetMysqlDB() (dbCon *gorm.DB) {
//	return db
//}
//
//func ConnectDatabase() {
//	issue := GetInstace().InitDataPool()
//	if !issue {
//		log.Println("Inital database pool fail")
//		os.Exit(1)
//	}
//	log.Println("Connect Database Success")
//}
