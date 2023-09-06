 URL Parsing

URLs provide a uniform way to locate resources. Here’s how to parse URLs in Go.
	


We’ll parse this example URL, which includes a scheme, authentication info, host, port, path, query params, and query fragment.
	


Accessing the scheme is straightforward.
	


User contains all authentication info; call Username and Password on this for individual values.
	



The Host contains both the hostname and the port, if present. Use SplitHostPort to extract them.
	


Here we extract the path and the fragment after the #.


To get query params in a string of k=v format, use RawQuery. You can also parse query params into a map. The parsed query param maps are from strings to slices of strings, so index into [0] if you only want the first value.
	


Running our URL parsing program shows all the different pieces that we extracted.