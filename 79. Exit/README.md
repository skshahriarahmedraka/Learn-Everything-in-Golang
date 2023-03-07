# Exit

#### Use os.Exit to immediately exit with a given status.

### Code

```
package main

import (
    "fmt"
    "os"
)

func main() {

    defer fmt.Println("!")

    os.Exit(3)
}
```

### Output

```
$ go run exit.go
exit status 3

$ go build exit.go
$ ./exit
$ echo $?
3
```

defers will not be run when using `os.Exit`, so this fmt.Println will never be called.

Exit with status 3.

Note that unlike e.g. C, Go does not use an integer return value from main to indicate exit status. If you’d like to exit with a non-zero status you should use os.Exit.

If you run exit.go using go run, the exit will be picked up by go and printed.

By building and executing a binary you can see the status in the terminal.
Note that the ! from our program never got printed.