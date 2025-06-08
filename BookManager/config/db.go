package config

import (
    "fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=localhost user=postgres password=yourpassword dbname=bookdb port=5432 sslmode=disable TimeZone=Asia/Amman"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    DB = db
    fmt.Println("Database connected")

    // تلقائيًا ينشئ جدول الكتب لو مش موجود
    DB.AutoMigrate(&Book{})
}
