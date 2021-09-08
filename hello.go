package main

import ("fmt"
		"math/rand"
		"strconv")

var msg = ` HELLO GOLANG!
  - coding
  - typing
  - exploring
`


func main() {

	fmt.Println(msg);

	str := "";
	num := 1

	for i := 0; i < 9; i++ {
		str += "hej: " + strconv.Itoa(num) + " \n"
		num += 1
	}

	fmt.Println(str)

	// for i := 0; i < 10; i++ {
	// 	// helloWorld(i)
	// 	// randomNumber(i)
	// 	x := "Hello! "
	// 	fmt.Printf("TYPE 1: %T  %d \n", i, i)
	// 	fmt.Printf("TYPE 2: %T %s \n", x, x)
	// 	even := i % 2 == 0
	// 	fmt.Printf("TYPE 3: %T %t \n", even, even)
	// }

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