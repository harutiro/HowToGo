package main 

import "fmt"

func main(){
	//スライスを作成する（関数によって）
	nums := make([]int,3,5)
	fmt.Println(len(nums),cap(nums))
	fmt.Println(nums)

	//len長さ　実際のよう素数
	//ｃap容量　その配列の最大の容量

	//長さの指定は省略できる
	nums2 := make([]int,0)
	fmt.Println(len(nums2),cap(nums2))
	fmt.Println(nums2)

	//要素を追加すると容量と長さが自動的に変化をする。
	nums3 := make([]int,0)
	fmt.Println(len(nums3),cap(nums3))
	nums3 = append(nums3,1)
	fmt.Println(len(nums3),cap(nums3))
	fmt.Println(nums3)

}