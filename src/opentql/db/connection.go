package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Connect() (dbobj gorm.DB) {

	dbobj, err := gorm.Open("mysql", "root:root@/idontknow?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	fmt.Println("Connection Successfull !")
	dbobj.SingularTable(true)

	return dbobj
}
