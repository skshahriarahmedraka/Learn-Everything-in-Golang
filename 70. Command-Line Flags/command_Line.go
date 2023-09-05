package main

import (
    "flag"
    "fmt"
)

func main() {

    wordPtr := flag.String("word", "foo", "a string")

    numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}

// $ go build command-line-flags.go

// $ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
// word: opt
// numb: 7
// fork: true
// svar: flag
// tail: []

	

// $ ./command-line-flags -word=opt
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: []

	

// $ ./command-line-flags -word=opt a1 a2 a3
// word: opt
// ...
// tail: [a1 a2 a3]

	

// $ ./command-line-flags -word=opt a1 a2 a3 -numb=7
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: [a1 a2 a3 -numb=7]

	

// $ ./command-line-flags -h
// Usage of ./command-line-flags:
//   -fork=false: a bool
//   -numb=42: an int
//   -svar="bar": a string var
//   -word="foo": a string

	

// $ ./command-line-flags -wat
// flag provided but not defined: -wat
// Usage of ./command-line-flags:
// ...