package main

import (
    "fmt"
    "time"
)

func main() {

    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}


// $ go run tickers.go
// Tick at 2012-09-23 11:29:56.487625 -0700 PDT
// Tick at 2012-09-23 11:29:56.988063 -0700 PDT
// Tick at 2012-09-23 11:29:57.488076 -0700 PDT
// Ticker stopped
