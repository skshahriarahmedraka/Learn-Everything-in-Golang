Parsing numbers from strings is a basic but common task in many programs; here’s how to do it in Go.
	


The built-in package strconv provides the number parsing.
	


With ParseFloat, this 64 tells how many bits of precision to parse.
	


For ParseInt, the 0 means infer the base from the string. 64 requires that the result fit in 64 bits.
	

ParseInt will recognize hex-formatted numbers.
	



A ParseUint is also available.
	



Atoi is a convenience function for basic base-10 int parsing.
	


Parse functions return an error on bad input.
	

