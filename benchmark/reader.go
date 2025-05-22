package benchmark

import (
    "github.com/yourusername/go-pg-bench/db"
    "gorm.io/gorm"
)

func ReadRecords(dbConn *gorm.DB, count int) {
    var recs []db.Record
    dbConn.Limit(count).Find(&recs)
}
