package persistance

import (
	"lsapp/model"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	// Register MySQL driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// Register database alias 'default'
	//orm.RegisterDataBase("default", "mysql", "root:(localhost:3306)/db")
	orm.RegisterDataBase("default", "mysql", "root:abc@123@tcp(localhost:3306)/ls?charset=utf8")
	// Set default database alias
	orm.SetDataBaseTZ("default", time.UTC)

	// Enable debug mode
	orm.Debug = true
	model.InitModel()
}
