package database

import (
    "os"

    "github.com/glebarez/sqlite"
    "gorm.io/gorm"
)

type Database interface {
    GetDB() *gorm.DB
}

type sqliteDB struct {
    db *gorm.DB
}

func (db *sqliteDB) GetDB() *gorm.DB {
    return db.db
}

func ConnectToDB() (Database, error) {
    dbPath := os.Getenv("SQLITE_DB")
    gormDB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return &sqliteDB{db: gormDB}, nil
}