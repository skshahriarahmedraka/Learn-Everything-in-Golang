Timers

We often want to execute Go code at some point in the future, or repeatedly at some interval. Go’s built-in timer and ticker features make both of these tasks easy. We’ll look first at timers and then at tickers.
	

Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides a channel that will be notified at that time. This timer will wait 2 seconds.
	
The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired.
If you just wanted to wait, you could have used time.Sleep. One reason a timer may be useful is that you can cancel the timer before it fires. Here’s an example of that.
	
Give the timer2 enough time to fire, if it ever was going to, to show it is in fact stopped.
	
The first timer will fire ~2s after we start the program, but the second should be stopped before it has a chance to fire.
