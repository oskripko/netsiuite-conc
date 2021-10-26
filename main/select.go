package main

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			println(x)
		case ch <- i:
		}
	}
}
