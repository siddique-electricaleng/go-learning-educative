package main

func main() {

	/*
		// TASK 1: Maps basics
		// Map literal
		mapLit := map[string]int{"one": 1, "two": 2}
		// Map assign
		var mapAssigned map[string]int

		// map creation using make()
		mapCreated := make(map[string]float32)

		mapAssigned = mapLit

		mapCreated["key1"] = 1.1
		mapCreated["key2"] = 9.1
		mapCreated["key3"] = 4.5

		fmt.Printf("Map literal:%v\n", mapLit)
		fmt.Printf("Map assigned:%v\n", mapAssigned)
		fmt.Printf("Map Created:%v\n", mapCreated)

		// Collection in the value of a key
		mapNewCollection := make(map[int]*[]int)

		slice1 := []int{10, 2, 3, 4}
		slice2 := []int{1111, 212, 33, 114}

		mapNewCollection[1] = &slice1
		mapNewCollection[2] = &slice2

		for key, val := range mapNewCollection {
			fmt.Printf("%d:%v\n", key, *val)
		}

		// Check if a key is present in a map using comma ok form
		value, isPresent := mapNewCollection[3]
		if isPresent {
			fmt.Printf("Key 3 is Present and has value:%v\n", value)
		} else {
			fmt.Printf("Key 3's presence is: %t\n", isPresent)
		}

		// Deleting an element with a key in maps
		delete(mapNewCollection, 2)

		for key, val := range mapNewCollection {
			fmt.Printf("%d : %v\n", key, *val)
		}

		// Slice of Maps
		// Version: A
		sliceOfMaps := make([]map[int]string, 4)

		for i := range sliceOfMaps {
			sliceOfMaps[i] = make(map[int]string, 1)
			sliceOfMaps[i][1] = strconv.Itoa(i) + "th Map"
			for key, val := range sliceOfMaps[i] {
				fmt.Printf("Key:%d -> Value:%s\n", key, val)
			}
		}
		// Version: B - not preferred it basically doesn't allocate any value at all

		items2 := make([]map[int]int, 5)
		for _, item := range items2 {
			item = make(map[int]int, 1) // item is only a copy of the slice element.
			item[1] = 2                 // This 'item' will be lost on the next iteration.
		}
		fmt.Printf("Version B: Value of items: %v\n", items2)
	*/
	// TASK LAST: Implement sorting and inverting a map
	sortInvertMain()
}
