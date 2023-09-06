package main

import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {

    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    r, _ := regexp.Compile("p([a-z]+)ch")

    fmt.Println(r.MatchString("peach"))

    fmt.Println(r.FindString("peach punch"))

    fmt.Println("idx:", r.FindStringIndex("peach punch"))

    fmt.Println(r.FindStringSubmatch("peach punch"))

    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    fmt.Println(r.FindAllString("peach punch pinch", -1))

    fmt.Println("all:", r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

    fmt.Println(r.FindAllString("peach punch pinch", 2))

    fmt.Println(r.Match([]byte("peach")))

    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println("regexp:", r)

    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}

 	

// $ go run regular-expressions.go
// true
// true
// peach
// idx: [0 5]
// [peach ea]
// [0 5 1 3]
// [peach punch pinch]
// all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
// [peach punch]
// true
// regexp: p([a-z]+)ch
// a <fruit>
// a PEACH