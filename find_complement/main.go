package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findComplement(5))
}

//转化成2进制字符串，每一个字符的整数都跟1异或，在转化为int
func findComplement(num int) int {
	numBinary := strconv.FormatInt(int64(num), 2)
	var conv []string
	numSlice := numBinary[0:]
	for _, value := range numSlice {
		temp := (value - 48) ^ 1
		tempStr := strconv.Itoa(int(temp))
		conv = append(conv, tempStr)
	}
	str := strings.Join(conv, "")
	res, _ := strconv.ParseInt(str, 2, 64)
	return int(res)
}
