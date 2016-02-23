package db

import (
	_"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	_"github.com/jinzhu/gorm"
	"opentql/db"
	_"time"
)

// Define model Structure

type account_info struct {
	ID           int    `gorm:primary;sql:"AUTO_INCREMENT"`
	Account_Name string `sql:"type:varchar(255)"`
	Cloud_Url    string `sql:"type:varchar(255)"`
	Username     string `sql:"type:varchar(255)"`
	Token        string `sql:"type:varchar(255)"`
	Site_Id      string `sql:"type:varchar(255)"`
}

type tql_master struct {
	ID               int    `gorm:primary;sql:"AUTO_INCREMENT"`
	Load_Balancer_Ip string `sql:"type:varchar(255)"`
	Site_Id          int
	Version          string `sql:"type:varchar(255)"`
	Status           string `sql:"type:varchar(255)"`
	Username         string `sql:"type:varchar(255)"`
	Password         string `sql:"type:varchar(255)"`
	Deployment_Type  string `sql:"type:varchar(255)"`
	Network_Name     string `sql:"type:varchar(255)"`
	Storage_Type     string `sql:"type:varchar(255)"`
	Image            string `sql:"type:varchar(255)"`
	Instance_Type    string `sql:"type:varchar(255)"`
	Facet_Key        string `sql:"type:varchar(255)"`
	Tql_Url          string `sql:"type:varchar(255)"`
}

type site struct {
	ID           int    `gorm:primary;sql:"AUTO_INCREMENT"`
	Name         string `sql:"type:varchar(255)"`
	Address      string `sql:"type:varchar(255)"`
	Geo_Location string `sql:"type:varchar(255)"`
	Pincode      string `sql:"type:varchar(255)"`
}

type gateway struct {
	Id          int    `gorm:primary;sql:"AUTO_INCREMENT"`
	Ip          string `sql:"type:varchar(255)"`
	Mac         string `sql:"type:varchar(255)"`
	Tql_Version string `sql:"type:varchar(255)"`
	Db_Name     string `sql:"type:varchar(255)"`
	Site_Id     int
	Status      string `sql:"type:varchar(255)"`
	Username    string `sql:"type:varchar(255)"`
	Password    string `sql:"type:varchar(255)"`
	Facet_Key   string `sql:"type:varchar(255)"`
	Snmp_Info   string `sql:"type:varchar(255)"`
}

type gateway_info struct {
	ID         int    `gorm:primary;sql:"AUTO_INCREMENT"`
	Ram        string `sql:"type:varchar(255)"`
	Version    string `sql:"type:varchar(255)"`
	Storage    string `sql:"type:varchar(255)"`
	Os         string `sql:"type:varchar(255)"`
	Gateway_Id int
	Cpu        string `sql:"type:varchar(255)"`
}

type metering struct {
	ID           int    `gorm:primary;sql:"AUTO_INCREMENT"`
	App_Events   string `sql:"type:varchar(255)"`
	Sensor_Event string `sql:"type:varchar(255)"`
	Gateway_Id   int
}

type monitoring struct {
	ID         int    `gorm:primary;sql:"AUTO_INCREMENT"`
	Watch_Dog  string `sql:"type:varchar(255)"`
	Network    string `sql:"type:varchar(255)"`
	Storage    string `sql:"type:varchar(255)"`
	Compute    string `sql:"type:varchar(255)"`
	Gateway_Id int
}


func CreateModels() {
	var dbobj = db.Connect()

	dbobj.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&account_info{})
	dbobj.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&tql_master{})
	dbobj.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&site{})
	dbobj.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&gateway{})
	dbobj.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&gateway_info{})
	dbobj.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&metering{})
	dbobj.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&monitoring{})

	fmt.Println("\n Tables Created")

	defer dbobj.Close()

}
