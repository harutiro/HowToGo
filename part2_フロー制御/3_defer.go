package main

import "fmt"

func main(){


	//処理をMain関数の最後まで遅延させることができる。
	defer fmt.Println("もうすぐ終わります")
	defer fmt.Println("ただいま勉強中！")
	fmt.Println("Go言語もやっています。")
}