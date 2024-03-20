package main

// Thread 1
func main() {
	// Thread 1 <-> Thread 2
	// The channel is a way to communicate between threads in Go language (goroutines).
	hello := make(chan string)

	// Thread 2
	go func() {
		hello <- "Hello, world!"
	}()

	result := <-hello
	println(result)
}
