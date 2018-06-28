package main

import (
	"fmt"

	"./funcpkg"
)

func main() {

	// nyuumonFunc() // 入門編
	// funcpkg.IfFunc() // if文
	// funcpkg.SwitchFunc("Java") // switch文
	// funcpkg.ForFunc() // for文
	// funcpkg.PointerFunc() // ポインタ
	// funcpkg.ReceiverFunc() // レシーバー
	// funcpkg.StructFunc() // 構造体
	// funcpkg.AboutFunc() // 関数いろいろ
	// funcpkg.InterfaceFunc() // インターフェース
	// funcpkg.SliceFunc() // スライス
	funcpkg.MapFunc()	// マップ
}

// 入門編
// ポイント
// 変数の宣言は型(string, intなど)を最後に持ってくる
func nyuumonFunc() {
	var str string     // 型を宣言
	str = "Start!!!!!" // 値を設定

	fmt.Println("Hellow Go World!! -> " + str)
}
