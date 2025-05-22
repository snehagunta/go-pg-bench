package util

import (
    "fmt"
    "time"
)

func Benchmark(taskName string, fn func()) {
    start := time.Now()
    fn()
    elapsed := time.Since(start)
    fmt.Printf("[%s] Took %s\n", taskName, elapsed)
}
