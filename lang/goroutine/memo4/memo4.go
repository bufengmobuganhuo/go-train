package memo4

import (
	"mengyu.com/gotrain/lang/goroutine/memo1"
)

// 代表一次函数调用请求
type request struct {
	Url string
	// 调用者需要的结果
	Response chan<- memo1.Result
}

type entry struct {
	res memo1.Result
	// 用于通知其他goroutine数据已准备好
	ready chan struct{}
}

func (e *entry) call(f memo1.Func, url string) {
	// 执行函数
	e.res.Value, e.res.Err = f(url)
	// 通知函数执行完毕
	close(e.ready)
}

func (e *entry) deliver(response chan<- memo1.Result) {
	// 等待数据准备完毕
	<-e.ready
	// 发送调用者需要的结果
	response <- e.res
}

type Memo struct {
	requests chan request
}

func (m *Memo) server(f memo1.Func) {
	// 缓存限制到一个方法内，线程安全
	cache := make(map[string]*entry)
	// 不停拉取客户端的调用请求
	for req := range m.requests {
		e := cache[req.Url]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.Url] = e
			// 真正调用这个方法
			go e.call(f, req.Url)
		}
		// 发送结果
		go e.deliver(req.Response)
	}
}

func New(f memo1.Func) *Memo {
	// 创建一个接收客户端调用请求的channel
	memo := &Memo{requests: make(chan request)}
	// 开始接收调用者的请求
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(url string) (interface{}, error) {
	// 用于接收结果的channel
	response := make(chan memo1.Result)
	memo.requests <- request{Url: url, Response: response}
	res := <-response
	return res.Value, res.Err
}
