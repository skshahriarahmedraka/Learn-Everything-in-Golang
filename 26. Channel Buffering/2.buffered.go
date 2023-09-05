package main

func main() {
	ch := make(chan int,1)
	ch<- 10  // non-blocking
	ChanMsg := <-ch  
	println(ChanMsg)

}