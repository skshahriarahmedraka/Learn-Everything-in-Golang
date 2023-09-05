
##  Spawning Processes

```
package main

import (
    "fmt"
    "io"
    "os/exec"
)

func main() {

    dateCmd := exec.Command("date")

    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

    _, err = exec.Command("date", "-x").Output()
    if err != nil {
        switch e := err.(type) {
        case *exec.Error:
            fmt.Println("failed executing:", err)
        case *exec.ExitError:
            fmt.Println("command exit rc =", e.ExitCode())
        default:
            panic(err)
        }
    }

    grepCmd := exec.Command("grep", "hello")

    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := io.ReadAll(grepOut)
    grepCmd.Wait()

    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}
```
### Output

```
$ go run spawning-processes.go 
> date
Thu 05 May 2022 10:10:12 PM PDT

date doesn’t have a -x flag so it will exit with an error message and non-zero return code.
	

command exited with rc = 1
> grep hello
hello grep

	

> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go

Next example: Exec'ing Processes. 
```

Sometimes our Go programs need to spawn other, non-Go processes.


	


We’ll start with a simple command that takes no arguments or input and just prints something to stdout. The exec.Command helper creates an object to represent this external process.


The Output method runs the command, waits for it to finish and collects its standard output. If there were no errors, dateOut will hold bytes with the date info.
	

Output and other methods of Command will return *exec.Error if there was a problem executing the command (e.g. wrong path), and *exec.ExitError if the command ran but exited with a non-zero return code.


Next we’ll look at a slightly more involved case where we pipe data to the external process on its stdin and collect the results from its stdout.
	


Here we explicitly grab input/output pipes, start the process, write some input to it, read the resulting output, and finally wait for the process to exit.


We omitted error checks in the above example, but you could use the usual if err != nil pattern for all of them. We also only collect the StdoutPipe results, but you could collect the StderrPipe in exactly the same way.
	


Note that when spawning commands we need to provide an explicitly delineated command and argument array, vs. being able to just pass in one command-line string. If you want to spawn a full command with a string, you can use bash’s -c option:
	



The spawned programs return output that is the same as if we had run them directly from the command-line.
	

date doesn’t have a -x flag so it will exit with an error message and non-zero return code.