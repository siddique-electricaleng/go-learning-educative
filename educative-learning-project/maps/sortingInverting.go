package main

import "fmt"

// Global Variables
var (
	sortBarVal = map[string]int{"kool": 12, "look": 11, "charlie": 100, "bob": 44, "phil": 30}

	invertBarVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87,
		"echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kilo": 43,
		"lima": 98}
)

func sortInvertMain() {

	// TASK 1: Sorting a map with its key,val pair
	// Sort that: map  -> slice -> sorted slice -> print the items from the map using
	// maps are unordered and therefore unsorted

	// TASK 1A: Print the unsorted map first with keys and values
	/*
		fmt.Println("unsorted:")
		fmt.Println("------------------------")
		fmt.Printf("%-10s | Value\n", "Key")
		fmt.Println("------------------------")
		for key, val := range sortBarVal {
			fmt.Printf("%-10s | %d\n", key, val)
		}
		// TASK 1B: Slice the map's keys out only
		keySlice := make([]string, len(sortBarVal))

		idx := 0
		for key := range sortBarVal {
			keySlice[idx] = key
			idx++
		}

		// TASK 2: String sort the keys inside the slice and print out final sorted map based on that
		sort.Strings(keySlice)
		fmt.Println("------------------------")
		fmt.Printf("%-20s\n", "Sorted Map")
		fmt.Println("------------------------")
		fmt.Printf("%-20s %0s %5s\n", "Key", "|", "Value")
		fmt.Println("------------------------")
		for _, val := range keySlice {
			fmt.Printf("%-20s %0s %5d\n", val, "|", sortBarVal[val])
		}
	*/

	// TASK 3: Invert a map

	/*
		Rules on inverting a map:
			- Inversion on a map === Swap keys & values
			- Key of original map and value of inverted map types must be same and vice versa
			- Map values also must be unique since they will be used as keys
	*/

	invertedMap := make(map[int]string, len(invertBarVal))

	// Simple as that, this is how inversion is done
	for key, val := range invertBarVal {
		invertedMap[val] = key
	}
	// Printing out the inverted map
	fmt.Println("Inverted Map")
	fmt.Println("------------------------")
	fmt.Printf("%-10s | Value\n", "Key")
	fmt.Println("------------------------")
	for key, val := range invertedMap {
		fmt.Printf("%-10d | %s\n", key, val)
	}

}
