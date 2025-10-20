package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
)

func InitDB() *gorm.DB {
    dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }
    // Auto-migrate models
    db.AutoMigrate(&models.Job{}, &models.Application{}, &models.User{})
    return db
}
