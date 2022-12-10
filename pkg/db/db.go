package db

import (
    "log"

    "github.com/Vfulgham/go-mux-api/pkg/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init() *gorm.DB {

    // connection string, really should be calling from .env file
    dbURL := "postgres://pg:pass@localhost:5432/crud"

    // open connection to Postgres db using dbURL
    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    // migrate models
    db.AutoMigrate(&models.Book{})

    return db
}
