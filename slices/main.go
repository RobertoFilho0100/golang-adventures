package main

import "fmt"

func main() {
	// Slice representating a shopping list
	shoppingList := []string{"Milk", "Bread", "Eggs"}

	// Adding items to the shopping List
	shoppingList = append(shoppingList, "Butter", "Coffee")

	// Updating an item
	shoppingList[1] = "Whole Gran Bread"

	// Displaying the full list
	fmt.Println("🛒 Shopping List:")
	for i, item := range shoppingList {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	// Creating a sublist with top items
	essentialItems := shoppingList[0:3] // First three items

	fmt.Println("\n✨ Essential Items:")
	for _, item := range essentialItems { // The underscore (_) is a blank identifier used to ignore the index
		fmt.Println("-", item)
	}

	// Checking the total number of items
	fmt.Printf("\n❓ Total items: %d\n", len(shoppingList))
}
