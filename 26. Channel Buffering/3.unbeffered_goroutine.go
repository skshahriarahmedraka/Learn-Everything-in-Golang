package main

// import "time"

func main() {
	ch := make(chan int)

	go func() {
		// time.Sleep(1 * time.Second)
		ChanMsg := <-ch //write to channel
		println(ChanMsg)

	}()
	ch<- 10 

}