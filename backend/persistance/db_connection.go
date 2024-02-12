package persistance

import (
	"fmt"
	"lsapp/model"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() error {
	// Register database alias 'default'
	dbURL := `root:Password@123@tcp(localhost:3306)/ls?charset=utf8`
	if err := orm.RegisterDataBase("default", "mysql", dbURL); err != nil {
		return fmt.Errorf("failed to register database: %w", err)
	}

	// Set default database alias
	orm.SetDataBaseTZ("default", time.UTC)

	// Enable debug mode
	orm.Debug = true

	// Initialize ORM models
	model.InitModel()

	return nil
}


