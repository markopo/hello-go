package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		helloWorld(i)
	}

}

func helloWorld(i int) {
	if(i %2 == 0) {
		fmt.Printf("hello, world: %d.\n", i)
	} else {
		fmt.Printf("Goodbye: %d \n", i)
	}
}