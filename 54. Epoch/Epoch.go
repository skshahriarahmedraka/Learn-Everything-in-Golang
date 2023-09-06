package main

import (
    "fmt"
    "time"
)

func main() {

    now := time.Now()
    fmt.Println(now)

    fmt.Println(now.Unix())
    fmt.Println(now.UnixMilli())
    fmt.Println(now.UnixNano())

    fmt.Println(time.Unix(now.Unix(), 0))
    fmt.Println(time.Unix(0, now.UnixNano()))
}

// $ go run epoch.go 
// 2012-10-31 16:13:58.292387 +0000 UTC
// 1351700038
// 1351700038292
// 1351700038292387000
// 2012-10-31 16:13:58 +0000 UTC
// 2012-10-31 16:13:58.292387 +0000 UTC

// Next we’ll look at another time-related task: time parsing and formatting.
	