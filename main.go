package main

import (
	"fmt"
	filter "project/bloom-filter"
)

func main() {
	fmt.Println("Implementations")
	bloomFilter := filter.NewBloomFilter()
	bloomFilter.Add("Akarshit")
	bloomFilter.Add("Jitesh")
	bloomFilter.Add("Tarun")
	bloomFilter.Add("Aman")

	if bloomFilter.IsPresent("Akarshi") {
		fmt.Println("Is present")
	} else {
		fmt.Println("Not Present")
	}

}
