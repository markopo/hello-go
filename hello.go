package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		helloWorld(i)
	}

}

func helloWorld(i int) {
	fmt.Printf("hello, world: %d.\n", i)
}