package database

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"member_church.com/infrastructure/entities"
)

func DatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	// 	DB_USER:="postgres"
	// DB_PASSWORD:="12345678"
	// DB_HOST:="127.0.0.1"
	// DB_NAME:="member_church_db"
	// DB_PORT:=5432
	// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	DB_USER,
	// 	DB_PASSWORD,
	// 	DB_HOST,
	// 	DB_PORT,
	// 	DB_NAME,
	// 	)
	dsn := "host=localhost user=postgres password=12345678 dbname=member_church_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	db.AutoMigrate(&entities.Rol{},
		&entities.Ministerial{},
		&entities.TeamPesca{},
		&entities.Church{},

		//&entities.UserMinisterial{},
	)
	return db
}

// Close database connection method is closin a connection between your app db
func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()

}

// Closedb
func Closedb() {
	var db *gorm.DB = DatabaseConnection()
	CloseConnection(db)

}
