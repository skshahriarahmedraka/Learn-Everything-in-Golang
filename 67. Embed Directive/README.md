
Go by Example: Embed Directive


Import the embed package; if you don’t use any exported identifiers from this package, you can do a blank import with _ "embed".
	


embed directives accept paths relative to the directory containing the Go source file. This directive embeds the contents of the file into the string variable immediately following it.
	

Or embed the contents of the file into a []byte.
	



We can also embed multiple files or even folders with wildcards. This uses a variable of the embed.FS type, which implements a simple virtual file system.
	


Print out the contents of single_file.txt.
	

Retrieve some files from the embedded folder.
	


Use these commands to run the example. (Note: due to limitation on go playground, this example can only be run on your local machine.)
