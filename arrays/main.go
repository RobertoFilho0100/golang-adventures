package main

import "fmt"

/*
This example demonstrates comprehensive array usage in Go
through a simple inventory and sales tracking system.

It covers:
- Declaration and initialization of arrays
- Accessing and updating elements
- Iterating with for-range loops
- Copying arrays and understanding value semantics
- Comparing arrays
- Multidimensional arrays to represent sales data per product and store location
- Basic calculation of totals and averages

Arrays are fixed-size collections holding elements of the same type.
*/

func main() {
	// Inventory stock levels for 4 products
	var stock [4]int = [4]int{100, 50, 75, 200}

	// Product names with inferred length
	products := [4]string{"Laptop", "Smartphone", "Tablet", "Headphones"}

	// Prices for each product
	prices := [4]float64{1500.00, 800.00, 450.00, 120.00}

	// Sales matrix: rows = products, columns = stores (3 stores)
	sales := [4][3]int{
		{30, 45, 25}, // Laptop sales in stores 1, 2, 3
		{50, 40, 60}, // Smartphone sales
		{25, 30, 20}, // Tablet sales
		{80, 90, 70}, // Headphones sales
	}

	// Print inventory and prices
	fmt.Println("Product Inventory and Prices:")
	for i, product := range products {
		fmt.Printf("%d. %s - Stock: %d units, Price: $%.2f\n", i+1, product, stock[i], prices[i])
	}

	// Update stock after some sales (for example: reduce Laptop stock by total sales)
	totalLaptopSales := 0
	for _, storeSales := range sales[0] {
		totalLaptopSales += storeSales
	}
	stock[0] -= totalLaptopSales

	fmt.Printf("\nUpdated stock for %s after sales: %d units\n", products[0], stock[0])

	// Calculate total sales per product across all stores
	fmt.Println("\nTotal sales per product:")
	for i, product := range products {
		totalSales := 0
		for _, storeSales := range sales[i] {
			totalSales += storeSales
		}
		fmt.Printf("%s: %d units sold\n", product, totalSales)
	}

	// Copy stock array to simulate a stock backup before changes
	backupStock := stock
	backupStock[1] = 100 // Modify backup stock of Smartphone

	fmt.Println("\nOriginal stock after backup modification:", stock)
	fmt.Println("Backup stock after modification:", backupStock)

	// Compare arrays
	areStocksEqual := stock == backupStock
	fmt.Printf("\nAre original stock and backup stock equal? %v\n", areStocksEqual)

	// Multidimensional array representing monthly sales for 3 months and 4 products
	monthlySales := [3][4]int{
		{75, 120, 80, 210}, // Month 1
		{60, 110, 75, 190}, // Month 2
		{90, 130, 95, 230}, // Month 3
	}

	fmt.Println("\nMonthly Sales per Product:")
	for monthIdx, month := range monthlySales {
		fmt.Printf("Month %d sales: ", monthIdx+1)
		for prodIdx, salesCount := range month {
			fmt.Printf("%s=%d ", products[prodIdx], salesCount)
		}
		fmt.Println()
	}

	// Calculate average sales per product over 3 months
	fmt.Println("\nAverage sales per product over 3 months:")
	for prodIdx := range products {
		total := 0
		for monthIdx := range monthlySales {
			total += monthlySales[monthIdx][prodIdx]
		}
		average := float64(total) / float64(len(monthlySales))
		fmt.Printf("%s: %.2f units\n", products[prodIdx], average)
	}

	/*
		Summary:
		- Arrays are used for fixed-size collections of related data.
		- Arrays are copied by value; modifying a copy does not affect the original.
		- Multidimensional arrays can represent tabular data, such as sales per store or per month.
		- Iteration allows access to all elements for calculation or display.
		- Array comparison checks length and each element for equality.
		- This example shows a practical inventory and sales tracking context.
	*/
}
