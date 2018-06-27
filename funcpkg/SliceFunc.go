package funcpkg

import (
	"fmt"
	"strconv"
)

// SliceFunc スライスのサンプルです
// ポイント
// 配列と似ているが値を参照型で保持している。
// スライス式"[:]"を使えばスライスになる
func SliceFunc() {
	val := [...]string{"A", "B", "C", "D"}

	fmt.Println("変更前 -> ", val)

	// スライスで関数に渡す（参照型なので値を変更できる)
	changeVal(val[:])

	fmt.Println("変更後 -> ", val)
}

func changeVal(valList []string) {
	for index := 0; index < len(valList); index++ {
		// スライスの値を変更
		valList[index] = valList[index] + strconv.Itoa(index+1)
	}
}
