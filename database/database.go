package database

import (
    "log"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/3x-haust/Go_TodoList/models"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect database")
    }

    err = DB.AutoMigrate(&models.Todo{})
    if err != nil {
        log.Fatal("Failed to migrate database")
    }
}