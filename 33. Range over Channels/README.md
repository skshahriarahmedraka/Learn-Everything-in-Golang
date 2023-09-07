Range over Channels

In a previous example we saw how for and range provide iteration over basic data structures. We can also use this syntax to iterate over values received from a channel.
	

We’ll iterate over 2 values in the queue channel.
	
This range iterates over each element as it’s received from queue. Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	
This example also showed that it’s possible to close a non-empty channel but still have the remaining values be received.
