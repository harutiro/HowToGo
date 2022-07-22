package main

import "fmt"

func main(){
	//Webのスライスができるよ
	names := []string{"yamada", "ikeda", "suzuki" , "tanaka"}

	fmt.Println(names)
	fmt.Println(names[0:4])
	fmt.Println(names[1:3])
	fmt.Println(names[2:])

	//可変長なので追加ができる
	fmt.Println("============================================================")
	names = append(names,"kato")

	fmt.Println(names)

}