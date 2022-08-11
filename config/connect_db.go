package config

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() *gorm.DB {
	DB = connectDB()
	return DB
}

func connectDB() *gorm.DB {
	dbName := os.Getenv("MYSQL_NAME")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbUser := os.Getenv("MYSQL_USERNAME")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}
	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	sqlDB, err := conn.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	return conn
}

func CloseConnectDB(conn *gorm.DB) {
	sqlDB, err := conn.DB()
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Close(); err != nil {
		panic(err)
	}
}
