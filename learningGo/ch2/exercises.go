package main

import "fmt"

const value = 100

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	var i int = 20
	var f float64 = float64(i)

	fmt.Println(i)
	fmt.Println(f)
}

func exercise2() {
	var i int = value
	var f float64 = value

	fmt.Println(i)
	fmt.Println(f)
}

func exercise3() {
	var b byte
	var smallI int32
	var bigI int64

	b = 255
	smallI = 2_147_483_647
	bigI = 9_223_372_036_854_775_807

	fmt.Println(b + 1)
	fmt.Println(smallI + 1)
	fmt.Println(bigI + 1)
}
