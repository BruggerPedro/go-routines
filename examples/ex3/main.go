package main

import "fmt"

func main() {

	forever := make(chan string)

	go func() {
		x := true
		for {
			if x == true {
				continue

			}
		}
	}()

	fmt.Println("Waiting forever")
	<-forever
	// While the forever channel is not closed, the program will not end.
	// The program will not end because the goroutine will be running forever and the main function will not be able to finish.
}
