package main

func main() {
	ch := make(chan int)
	ch<- 10  //blocking
	ChanMsg := <-ch //write to channel
	println(ChanMsg)

}