package main

import "fmt"

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Printf("\n")

}

//variadic function
func sum(nums ...int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

// Creating a matrix whre matrix[i][j] = i* j
func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)

	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	// Arrays have fixed size
	var arr [5]int
	printArr(arr[:])

	// Slices are built on top of arrays
	slice := []int{1, 2, 3, 4, 5}
	printArr(slice)

	// create slices with make, value for int defaults to 0
	makeSlice := make([]int, 10)
	printArr(makeSlice)
	fmt.Printf("max cap of slice is %v\n", cap(makeSlice))
	fmt.Printf("len of slice is %v\n", len(makeSlice))

	// Calling variadic functions
	fmt.Printf("Calling sum(1,2,3) \n")
	fmt.Printf("Result is %d", sum(1, 2, 3))

	// Spread operator
	numbers := []int{1, 2, 3, 4, 5}
	total := sum(numbers...)
	fmt.Printf("Result of using spread operator for numbers 1,2,3,4,5 is %d \n", total)

	// append function
	numbers = append(numbers, 1)
	numbers = append(numbers, 2, 3)
	fmt.Printf("Result of using spread operator for numbers 1,2,3,4,5,1,2,3 is %d\n", sum(numbers...))

	newMatrix := createMatrix(5, 5)
	fmt.Printf("Printing New Matrix \n")
	printMatrix(newMatrix)
}
