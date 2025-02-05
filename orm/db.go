package orm

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

// User model (Make sure this struct exists)
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Fullname string
	Avatar   string
}

func InitDB() {
	var err error
	dsn := "root:SuriyaMysql107@tcp(127.0.0.1:3306)/go-jwt?charset=utf8mb4&parseTime=True&loc=Local"

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	log.Println("✅ Database connected successfully")

	// Migrate the User model
	Db.AutoMigrate(&User{})
}
