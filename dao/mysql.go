package dao

import (
	"gorm.io/gorm"
	"fmt"
	"gorm.io/gorm/schema"
	"gorm.io/driver/mysql"
	"awesomeProject7/module"
)

var DB *gorm.DB

type Config struct {
	user   string
	pass   string
	adrr   string
	port   string
	dbName string
	time   string
}

func ConnDB() *gorm.DB {
	conf := &Config{"root", "root", "localhost", "3306", "ljj", "5s"}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", conf.user, conf.pass, conf.adrr, conf.port, conf.dbName, conf.time)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("connect failed, error=" + err.Error())
	}
	db.Create(&module.TaskResult{})
	db.AutoMigrate(&module.TaskResult{})
	DB = db
	return db
}

func UpgradeResult(result string, tasknote string) {
	taskResult := module.TaskResult{
		TaskNote: tasknote,
	}
	DB.Model(&taskResult).Updates(map[string]interface{}{"Result": result})
}
