Go by Example: HTTP Client

The Go standard library comes with excellent support for HTTP clients and servers in the net/http package. In this example we’ll use it to issue simple HTTP requests.


Issue an HTTP GET request to a server. http.Get is a convenient shortcut around creating an http.Client object and calling its Get method; it uses the http.DefaultClient object which has useful default settings.
	



Print the HTTP response status.
	

Print the first 5 lines of the response body.
