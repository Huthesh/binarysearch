package main 

import (
	"binarysearch/algo"
	"fmt"
)

func main() {
	array:=[]int{2, 3, 5, 7, 11, 13}
	
	fmt.Println(algo.BinarySearch(array,7))
	
	fmt.Println(algo.BinarySearch(array,1))
	
	fmt.Println(algo.BinarySearch(array,13))
	
	
	fmt.Println(algo.BinarySearch(array,4))
}
