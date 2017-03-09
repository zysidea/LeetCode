### Hamming Distance
先将两个数异或，然后循环判断每一个bit是否为1
```go
func hammingDistance(x int, y int) int {
	res:= 0
	exc:= x ^ y
	for i:= 0; i < 32;i++  {
		j:=uint(i)
		res += (exc >> j) & 1
	}
	return res
}
```
### Number Complement 
转化成2进制字符串，每一个字符的整数都跟1异或，在转化为int
```go
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
```
