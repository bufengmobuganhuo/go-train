package main

import "fmt"

// 字符串的最长不重复子串
func lengthOfNonRepeatingSubStr(s string) int {
	// 字符最后一次出现的位置
	lastOccurred := make(map[byte]int)
	startPtr := 0
	maxLen := 0
	// 这里用的byte，故只适用于ASCII码
	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		// 如果字符ch最后一次出现的位置比startPtr大，则更新窗口
		if ok && lastI >= startPtr {
			startPtr = lastOccurred[ch] + 1
		}
		if i-startPtr+1 > maxLen {
			maxLen = i - startPtr + 1
		}
		lastOccurred[ch] = i
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcab"))
}
