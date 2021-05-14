package main

import "fmt"

// TIDAK BISA KARNA MAIN HARUS TIDAK MEMILIKI ARGUMENTS
func main(nums ...int) {

	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}

main(1, 2)
main(1, 2, 3)

nums := []int{1, 2, 3, 4}
main(nums...)