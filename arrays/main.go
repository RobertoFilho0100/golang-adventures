package main

import "fmt"

/*
This example demonstrates the use of arrays in Go.
Arrays have a fixed size and store elements of the same type.
Unlike slices, their length is part of their type.

The code covers:
- Declaration and initialization
- Accessing and updating elements
- Iteration
- Copying arrays
- Comparison
- Multidimensional arrays
*/
func main() {
	// Declaring and initializing an array of 5 integers
	var scores [5]int = [5]int{90, 85, 78, 92, 88}

	// Alternatively, let Go infer the length
	colors := [3]string{"Red", "Green", "Blue"}

	// Accessing elements by index (zero-based)
	fmt.Println("First score:", scores[0])
	fmt.Println("Second color:", colors[1])

	// Updating an element
	scores[2] = 80
	fmt.Println("Updated scores:", scores)

	// Iterating over the array with index and value
	fmt.Println("\nScores list:")
	for i, score := range scores {
		fmt.Printf("Score %d: %d\n", i+1, score)
	}

	// Copying arrays creates a new, independent copy (unlike slices)
	scoresCopy := scores
	scoresCopy[0] = 100 // Modifying copy does not affect original
	fmt.Println("\nOriginal scores after copy modification:", scores)
	fmt.Println("Modified copy of scores:", scoresCopy)

	// Comparing arrays (must have the same length and element type)
	equal := scores == scoresCopy
	fmt.Printf("\nAre scores and scoresCopy equal? %v\n", equal)

	// Multidimensional arrays (e.g., a 3x3 matrix)
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("\nMatrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	/*
		Summary:
		- Arrays have fixed size defined at compile-time.
		- They hold elements of the same type.
		- Copying arrays copies all elements.
		- Use arrays when fixed size and memory layout is important.
		- For flexible, resizable collections, slices are preferred.
	*/
}
