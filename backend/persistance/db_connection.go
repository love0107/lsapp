package persistance

import (
	"fmt"
	"lsapp/model"
	"os"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Init() error {
	// password := os.Getenv("DB_PASSWORD")
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed to load env variables: %w", err)
	}
	password := os.Getenv("DB_PASSWORD")
	dbString := `user=postgres.kxzrdckifcablkhdludv password=` + password + ` host=aws-0-ap-south-1.pooler.supabase.com port=5432 dbname=postgres`
	// Register database alias 'default'
	if err := orm.RegisterDataBase("default", "postgres", dbString); err != nil {
		return fmt.Errorf("failed to register database: %w", err)
	}

	// Set default database alias timezone
	orm.SetDataBaseTZ("default", time.UTC)

	// Enable debug mode
	orm.Debug = true

	// Initialize ORM models
	model.InitModel()

	return nil
}
