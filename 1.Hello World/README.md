Hello World

Our first program will print the classic “hello world” message. Here’s the full source code.

## hello-world.go
```
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```
## Output
```
$ go run hello-world.go
hello world


$ go build hello-world.go
$ ls
hello-world    hello-world.go

$ ./hello-world
hello world
```
To run the program, put the code in hello-world.go and use go run.
	
Sometimes we’ll want to build our programs into binaries. We can do this using go build.
	
We can then execute the built binary directly.
	


Now that we can run and build basic Go programs, let’s learn more about the language.