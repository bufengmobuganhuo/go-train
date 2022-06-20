package thumbnail

import (
	"log"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
)



// 模式4：使用sync.WaitGroup等待一些goroutine结束
func makeThumbnails4(filenames []string) int64 {
	sizes := make(chan int64)
	// 正在工作中的goroutine的个数
	var wg sync.WaitGroup
	for _, f := range filenames {
		wg.Add(1)
		go func(f string) {
			// 保证一定可以让wg -1
			defer wg.Done()
			thumbnail, err := thumbnail.ImageFile(f)
			// 如果发生错误，则直接结束
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumbnail)
			// 获取压缩后文件的大小
			sizes <- info.Size()
		}(f)
	}

	// 启动一个goroutine，等待上面的任务结束。
	// 这里如果不使用goroutine，那么会一直阻塞，那么会造成下面的代码一直不可达
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

// 模式3：同模式2，将详细信息返回给外部goroutine
func makeThumbnails3(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	// 构建缓冲通道
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		// 使用入参的方式，也可以防止迭代错误
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			// 将详细信息传送到共享通道
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		// 这里遇到第一个错误时就返回
		// 如果上面没有使用缓冲通道，
		// 那么对于其他goroutine，发送操作会一直阻塞，从而导致goroutine无法被回收而泄露
		if it.err != nil {
			return thumbfiles, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

// 模式2：使用一个共享的通道发送时间来向外层goroutine报告他已经完成
func makeThumbnails2(filenames []string) {
	ch := make(chan struct{})

	for _, f := range filenames {
		// 变量f的作用域是整个循环体，在循环里创建的所有函数变量都会共享相同的变量
		//（传入给thumbnail.ImageFile()函数的是同一个变量）（他是一个可访问的地址，而不是固定的值），
		// 他的值在迭代中会被不断更新，因此在循环结束时f的值是最后一次迭代时的值
		file := f
		go func() {
			thumbnail.ImageFile(file)
			// 发送任务执行完成的事件
			ch <- struct{}{}
		}()
	}

	// 外部goroutine等待任务完成
	for range filenames {
		<-ch
	}
}

// 模式1：直接启动goroutine，不在乎goroutine是否正常结束
func makeThumbnails1(filenames []string) {
	for _, f := range filenames {
		file := f
		go thumbnail.ImageFile(file)
	}
}
