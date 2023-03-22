package initializers

import (
	"fmt"

	"github.com/keima483/fiber-ems/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitialiseDB(config *Config) error {
	DNS := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		config.DBUserName, 
		config.DBUserPassword, 
		config.DBHost, 
		config.DBPort, 
		config.DBName,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("DB Connection failed")
		return err
	}
	return DB.AutoMigrate(&repository.Company{}, &repository.Employee{})
}
