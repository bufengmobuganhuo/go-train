package nodes

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

func Init() {
	startTime = time.Now()
}

func ArraySource(a ...int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		// 送完后需要关闭channel
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		fmt.Println("Read done:", time.Since(startTime))
		sort.Ints(a)
		fmt.Println("InMemSort done:", time.Since(startTime))
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	// 给定初始值大小，可以让从channel获取元素的阻塞时间减少
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		// 归并排序思想，从两个channel中分别取一个，取较小值
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("Merge done:", time.Since(startTime))
	}()
	return out
}

// 读取数据来源
func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		// 8*8=64 bit的slice，因为go语言的int随系统而变化，本机是64位系统
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				// 转化成int
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			// 分块读，超过chunk的大小，则停止读取
			if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

// 将结果写入某个位置
func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	// 归并
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))
}
