package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/Users/tiger/err.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	map1 := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if io.EOF == err {
			break
		}
		if err != nil {
			panic(err)
		}
		idx := strings.Index(line, "planId=")
		if idx < 0 {
			continue
		}
		var planId bytes.Buffer
		idx += 7
		for ; line[idx] != ','; idx++ {
			planId.WriteByte(line[idx])
		}
		idx += 14
		var algoOrderId bytes.Buffer
		for ; line[idx] != ','; idx++ {
			algoOrderId.WriteByte(line[idx])
		}
		map1[planId.String()] = algoOrderId.String()
	}
	fmt.Printf("len=%d\n", len(map1))
	for k, v := range map1 {
		fmt.Printf("planId=%s, algoOrderId=%s\n", k, v)
	}
}
