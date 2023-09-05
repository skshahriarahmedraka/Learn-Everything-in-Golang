Go by Example: Command-Line Arguments

Command-line arguments are a common way to parameterize execution of programs. For example, go run hello.go uses run and hello.go arguments to the go program.
	
	[Run code] [Copy code]

package main

	


os.Args provides access to raw command-line arguments. Note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program.
	


You can get individual args with normal indexing.
	



To experiment with command-line arguments it’s best to build a binary with go build first.
	



Next we’ll look at more advanced command-line processing with flags