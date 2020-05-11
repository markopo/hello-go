package main

import ("fmt"
		"math/rand")

func main() {

	for i := 0; i < 10; i++ {
		helloWorld(i)
		randomNumber(i)
	}

}

func randomNumber(i int) {
	r := rand.Intn(100)
	sum := i * r
	_, e := fmt.Printf("Random number: %d : %d : %d \n", i, r, sum)

	if (e != nil) {
		fmt.Println(e)
	}

}

func helloWorld(i int) {
	if(i %2 == 0) {
		fmt.Printf("hello, world: %d.\n", i)
	} else {
		fmt.Printf("Goodbye: %d \n", i)
	}
}