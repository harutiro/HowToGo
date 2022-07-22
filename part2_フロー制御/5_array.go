package main

import "fmt"

func main(){

	//スライスを取り出す
	numArray := [5]int{3,6,9,12,15}
	for i , v := range numArray{
		fmt.Println(i,v)
	}

}