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
