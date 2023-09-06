Epoch

A common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds since the Unix epoch. Here’s how to do it in Go.
	

Use time.Now with Unix, UnixMilli or UnixNano to get elapsed time since the Unix epoch in seconds, milliseconds or nanoseconds, respectively.


You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.
	

	

