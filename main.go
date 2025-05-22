package main

import (
    "github.com/yourusername/go-pg-bench/config"
    "github.com/yourusername/go-pg-bench/db"
    "github.com/yourusername/go-pg-bench/benchmark"
    "github.com/yourusername/go-pg-bench/util"
)

func main() {
    dbConn := config.ConnectDB()
    dbConn.AutoMigrate(&db.Record{})

    util.Benchmark("Write 1000 records", func() {
        benchmark.WriteRecords(dbConn, 1000)
    })

    util.Benchmark("Read 1000 records", func() {
        benchmark.ReadRecords(dbConn, 1000)
    })
}
