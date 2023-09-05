Channel Buffering

By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.

```
package main

import "fmt"

func main() {

    messages := make(chan string, 2)

    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
```


Here we make a channel of strings buffering up to 2 values.
	

Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	

Later we can receive these two values as usual

//////////////////////////////
If the channel is unbuffered (capacity is zero), then communication succeeds only when the sender and receiver are both ready.

If the channel is buffered (capacity >= 1), then send succeeds without blocking if the channel is not full and receive succeeds without blocking if the buffer is not empty.

    When putting a value to the intChannelZero like intChannelZero <- 1, where the value be saved?

The value is copied from the sender to the receiver. The value is not saved anywhere other than whatever temporary variables the implementation might use.

    The differences between intChannelZero and intChannelOne when putting a value to them.

Send on intChannelZero blocks until a receiver is ready.

Send on intChannelOne blocks until space is available in the buffer.



