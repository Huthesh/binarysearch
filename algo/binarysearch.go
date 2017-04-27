package algo

import (

)

func BinarySearch(array []int,key int) (result bool){	
		startIndex:=0
		length:=len(array)
		lastIndex:=length-1
		for startIndex<=lastIndex {
			mid:=(startIndex+lastIndex)/2
			switch{
				case array[mid]==key :
					return true
				case key>array[mid] :
					startIndex=mid+1
				case key<array[mid]:
					lastIndex=mid-1
			}
		}
		return false
}