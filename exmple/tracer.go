// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("test")
// 	res := Sum(1, 2, 3, 4)
// 	fmt.Println(res)
// }

// // write a sum number which is going to accept a range of numbers
// func Sum(numbers ...int) int {
// 	total := 0
// 	for _, n := range numbers {
// 		total += n
// 	}
// 	return total
// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("test")
// 	a := []int{1, 2, 3, 4, 5}
// 	fmt.Println(twoSums(a, 7))
// }

// func twoSums(nums []int, target int) []int {
// 	mymap := make(map[int]int)
// 	for i := 0; i < len(nums); i++ {
// 		j, ok := mymap[target-nums[i]]
// 		if ok {
// 			result := []int{j, i}
// 			return result
// 		}
// 		mymap[nums[i]] = i
// 	}
// 	result := []int{-1, 1}
// 	return result

// }
