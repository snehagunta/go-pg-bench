package benchmark

import (
    "github.com/yourusername/go-pg-bench/db"
    "gorm.io/gorm"
    "math/rand"
)

func WriteRecords(dbConn *gorm.DB, count int) {
    for i := 0; i < count; i++ {
        dbConn.Create(&db.Record{
            Name:  "record",
            Value: rand.Intn(100),
        })
    }
}
