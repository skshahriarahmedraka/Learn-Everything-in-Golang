package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    now := time.Now()
    p(now)

    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    p(then.Weekday())

    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    diff := now.Sub(then)
    p(diff)

    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    p(then.Add(diff))
    p(then.Add(-diff))
}

 	

// $ go run time.go
// 2012-10-31 15:50:13.793654 +0000 UTC
// 2009-11-17 20:34:58.651387237 +0000 UTC
// 2009
// November
// 17
// 20
// 34
// 58
// 651387237
// UTC
// Tuesday
// true
// false
// false
// 25891h15m15.142266763s
// 25891.25420618521
// 1.5534752523711128e+06
// 9.320851514226677e+07
// 93208515142266763
// 2012-10-31 15:50:13.793654 +0000 UTC
// 2006-12-05 01:19:43.509120474 +0000 UTC