package main

import (
	"fmt"
)

func intToByteArray(num int) []byte {
	// 将数字转换为字符串
	numStr := fmt.Sprintf("%d", num)
	
	// 创建一个 []byte 数组
	result := make([]byte, len(numStr))

	// 将字符串中的每个字符存储到数组中
	for i, char := range numStr {
		result[i] = byte(char)
	}

	return result
}

func main() {
	num := 32

	// 将整数按位存储到 []byte 数组中
	byteArray := intToByteArray(num)
	j := 99
	numStr := fmt.Sprintf("%d", j)
	fmt.Println(numStr)
	fmt.Printf("%T" , numStr)
	// 输出结果
	fmt.Println(byteArray)
}