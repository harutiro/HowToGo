package main

import "fmt"

func main(){
	lang := "go"

	switch lang{
	case "go":
		fmt.Println("Go")
	case "java":
		fmt.Println("Java")
	case "python":
		fmt.Println("Python")
	default:
		fmt.Println("Unknown")
	}
}