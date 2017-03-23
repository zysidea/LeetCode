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
### Keyboard Row
题目中没有指定键盘上的字母是否是大小写，因此对于每一行都定义为一个slice(这个办法有点low，但是简单的设计貌似并没有什么规律)，然后进行循环去判断单词是否包含一行中的任意一个字母，只有只匹配一行的时候才可以打印该字符串
```go
var (
	firstKey = []string{"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P", "q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}

	secondKey = []string{"A", "S", "D", "F", "G", "H", "J", "K", "L", "a", "s", "d", "f", "g", "h", "j", "k", "l"}

	thirdKey = []string{"Z", "X", "C", "V", "B", "N", "M", "z", "x", "c", "v", "b", "n", "m"}
)

func findWords(words []string) []string {
	var result []string
	for _, v1 := range words {
		match1 := match(v1, firstKey)
		match2 := match(v1, secondKey)
		match3 := match(v1, thirdKey)

		if match1 && !match2 && !match3 {
			result = append(result, v1)
		} else if !match1 && match2 && !match3 {
			result = append(result, v1)
		} else if !match1 && !match2 && match3 {
			result = append(result, v1)
		} else {
			continue
		}
	}
	return result
}

func match(str string, matchs []string) bool {
	for _, v := range matchs {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}
```
