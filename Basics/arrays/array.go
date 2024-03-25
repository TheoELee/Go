package main

import (
	"fmt"
	"strings"
)

func main() {
	basicArray()
	slices()
}

func basicArray() {
	// the type [n]T is an array of n values of type T
	var a [2]string
	a[0] = "Hello"
	a[1] = "World!"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	fmt.Println("\n--- SLICES ---")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// arrays have a fixed sizes. We can use "slice", a dynamically sized view into the array, and is much more common
	// []T is a slice with elements of type T. Slice is formed by specifying a low and high index a[log : high]
	// note, does not include the last index
	var slice []int = primes[1 : 4]
	fmt.Println(slice)

	// slices do not store any data, they only describe a section of the array.
	// !! changing the elements of a slice modifies the corresponding elements of the underlying array
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	// once the value is changed, all references to that underlying array reflect the changes
	b[0] = "XXXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// SLICE LITERALS
	// like an array literal, without the length
	arrayLiteral := [3]bool{true, true, false}
	sliceLiteral := []bool{true, true, false}
	fmt.Println(arrayLiteral, sliceLiteral)

	// SLICE DEFAULTS
	// low bound default is zero
	// high bound default is array.length
	// the bounds can be omitted if you want to just use the defaults
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// these are all equal expressions
	// array[0 : 10]
	// array[:10]
	// array[0:]
	fmt.Println("slice using defaults", array[:])

	// SLICE LENGTH AND CAPACITY
	// length of the slice is the num elements it contains
		// you can extend a slices length by re-slicing, assuming it has sufficient capacity 
	// capacity is the num elements in the underlying array
	sliceLengthAndCapacity := func() {
	fmt.Println("\n~ Slice length and capacity ~")
		printSlice := func (s []int) {
			fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
		}

		s := []int{2, 3, 5, 7, 11, 13}
		printSlice(s)

		// Slice the slice to give it zero length
		s = s[:0]
		printSlice(s)

		// Extend its length
		s = s[:4]
		printSlice(s)

		// Drop the first two values
		s = s[2:]
		printSlice(s)
	}
	sliceLengthAndCapacity()

	// NIL SLICES
	// the zero value of a slice is nil. A nil slice has length/capacity of zero
	nilSlice := func() {
	fmt.Println("\n~ Nil slice ~")
		var s []int
		fmt.Println(s, len(s), cap(s))

		if s == nil {
			fmt.Println("nil!")
		}
	}
	nilSlice()

	// SLICE AND "MAKE"
	// slices can be created with the built in `make` function, this is how dynamically sized 
	// array are created
	// `make` allocates a zeroes array and returns a slice that refers to that array 
	creatingSliceWithMake := func() {
		fmt.Println("\n~ Create slice with `make` ~")
		printSlice := func(s string, x []int) {
			fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
		}

		a := make([]int, 5)
		printSlice("a", a)

		b := make([]int, 0, 5)
		printSlice("b", b)

		c := b[:2]
		printSlice("c", c)

		d := c[2 : 5]
		printSlice("d", d)
	}
	creatingSliceWithMake()

	// SLICES OF SLICES (that's enough slices!)
	// slices can contain any type, including other slices
	slicesOfSlices := func() {
		fmt.Println("\n~ Slices of slices ~")
		// tic-tac-toe board
		board := [][]string {
			[]string{"_", "_", "_", "_"},
			[]string{"_", "_", "_", "_"},
			[]string{"_", "_", "_", "_"},
		}

		// players take turns
		board[0][0] = "X"
		board[2][2] = "O"
		board[1][2] = "X"
		board[1][0] = "O"
		board[0][2] = "X"

		for i := 0; i < len(board); i++ {
			fmt.Printf("%s\n", strings.Join(board[i], " "))
		}
	}
	slicesOfSlices()
}