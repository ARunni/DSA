// Remove near by duplicates in an array
package main

import "fmt"


func removeDuplicates(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var newArray []int
	newArray = append(newArray, arr[0])
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			newArray = append(newArray, arr[i])
		}		
	}
	copy(arr,newArray)
	length := len(newArray)
	return length
}

func main()  {
	arr := []int{1,1,2,2,3,3,4,5}
	fmt.Println("Before Removing elements",arr)
	newLength := removeDuplicates(arr)
	fmt.Println("After",arr[:newLength])

}
