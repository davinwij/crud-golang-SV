package connection

import (
	"fmt"
	"log"
	"simple-crud-task/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB  *gorm.DB
	Err error
)

func ConnectDB() {

	DB, Err = gorm.Open("mysql", "root:@/simple-crud-golang?charset=utf8&parseTime=True")

	if Err != nil {
		log.Fatal(Err)
	}

	fmt.Println("Database Connection Success")

	DB.AutoMigrate(models.USER{})
	DB.AutoMigrate(models.RISK_PROFILE{})
}
