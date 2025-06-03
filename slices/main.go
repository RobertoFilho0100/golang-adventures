package main

import "fmt"

/*
This example demonstrates various operations with slices in Go,
including creation, updating, appending, sub-slicing, capacity inspection,
memory sharing, and element removal.
*/
func main() {
	/*
		Creating a slice with initial values.
		A slice is a dynamically-sized, flexible view into the elements of an array.
	*/
	shoppingList := []string{"Milk", "Bread", "Eggs"}

	// Adding items using append.
	shoppingList = append(shoppingList, "Butter", "Coffee")

	// Updating an item by index.
	shoppingList[1] = "Whole Grain Bread"

	// Displaying the full shopping list.
	fmt.Println("Shopping List:")
	for i, item := range shoppingList {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	// Checking length and capacity.
	fmt.Printf("\nLength: %d, Capacity: %d\n", len(shoppingList), cap(shoppingList))

	/*
		Creating a sub-slice that includes the first three items.
		This sub-slice shares the same underlying array as the original slice.
	*/
	essentialItems := shoppingList[0:3]

	fmt.Println("\nEssential Items:")
	for _, item := range essentialItems { // The underscore (_) is called the blank identifier. It is used to ignore the index returned by range when the index is not needed.
		fmt.Println("-", item)
	}

	/*
		Demonstrating memory sharing between slices.
		Updating an item in the sub-slice also affects the original slice
		because both slices point to the same underlying array.
	*/
	essentialItems[0] = "Water"

	fmt.Println("\nAfter updating essentialItems[0] to 'Water':")
	fmt.Println("Essential Items:", essentialItems)
	fmt.Println("Shopping List:", shoppingList)

	/*
		Creating a slice using make().
		This creates a slice with length 0 but capacity for 5 elements.
		Useful when the number of items is unknown but the capacity is expected to grow.
	*/
	tasks := make([]string, 0, 5)
	tasks = append(tasks, "Buy milk", "Clean house")

	fmt.Println("\nTasks List:", tasks)
	fmt.Printf("Tasks - Length: %d, Capacity: %d\n", len(tasks), cap(tasks))

	/*
		Removing an item from a slice.
		For example, removing the item at index 2 ("Eggs").
		Steps:
		- Take all elements before the index.
		- Append all elements after the index.
	*/
	indexToRemove := 2
	shoppingList = append(shoppingList[:indexToRemove], shoppingList[indexToRemove+1:]...)

	fmt.Println("\nShopping List after removing item at index 2 (Eggs):")
	for i, item := range shoppingList {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}
