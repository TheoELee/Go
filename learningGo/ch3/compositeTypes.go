package main

import "fmt"
import "slices"

func main() {
	arrays()
}

func arrays() {
	// arrays are rarely used directly in Go. But here's how they're declared

	// note: all elements in the array must be of the type specified
	var baseArray [3]int // an array of 3 ints (no values were specified so all entires are set to zero)

	var valueArray = [3]int{10, 20, 30} // an array 3 ints with values specified

	var sparseArray = [5]int{1, 3: 4} // can set value of specific index - {1, 0, 0, 4, 0}

	var withArrayLiteral = [...]int{10, 20, 30} // don't need to define type when insantiating from array literal

	// as with other lower level languages, you cannot write past the end of an array or use a negative index
	// doing so results in a compile time error

	// built in function len returns length of array

	fmt.Println(baseArray)
	fmt.Println(valueArray)
	fmt.Println(sparseArray)
	fmt.Println(withArrayLiteral)

	// arrays in Go are unique in that the size of the array is part of the type of the array.
	// meaning that, an array of [3]int is a different type altogether from an array of [4]int. This
	// also means that a variable can't be used to specificy the size of an array since the size needs
	// to be determined at compile time

	// furthermore, you ** cannot use type conversion to convert arrays of different sizes to identical types **.
	// Meaning that you can't write a function that works with arrays of any size and you can't assign arrays
	// of different sizes of the same variable
}

func slices_ch3() {
	// Slices are growable arrays and much more common than arrays in Go
	// One of the key differences is that the length IS NOT part of the type of a slice

	// Defining a slice:
	var x = []int{10, 20, 30}
	var y = []int{1, 5: 4, 7: 3}

	// Multi-dimensional
	var multi [][]int

	// Zero value slice
	var zeroSlice []int

	// Note, slices ARE NOT comparable. It's a compile time error to check if two slices are
	// identical or not with != or ==. Can only compare to nil value. Instead, standard lib includes
	// two functions to compare slices:
	slices.Equal(x, y)

	slices.EqualFunc(x, test)

}
