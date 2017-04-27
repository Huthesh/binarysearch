package algo

import (

)

func BinarySearch(array []int,key int) (result bool){	
		startIndex:=0
		length:=len(array)
		lastIndex:=length-1
		for startIndex<=lastIndex {
			mid:=(startIndex+lastIndex)/2
			if(array[mid]==key){
				return true;
			}
			if(key>array[mid]){
				startIndex=mid+1
			} else if(key<array[mid]) {
				lastIndex=mid-1
			}
			
		}
		return false
}