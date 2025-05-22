package db

import "time"

type Record struct {
    ID        uint `gorm:"primaryKey"`
    Name      string
    Value     int
    CreatedAt time.Time
}
