package memo2

import (
	"sync"

	"mengyu.com/gotrain/lang/goroutine/memo1"
)

// 使用同步锁实现并发安全
type Memo struct {
	f     memo1.Func
	cache map[string]memo1.Result
	mu    sync.Mutex
}

func New(f memo1.Func) *Memo {
	return &Memo{f: f, cache: make(map[string]memo1.Result)}
}

func (m *Memo) Get(url string) (interface{}, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	res, ok := m.cache[url]
	if !ok {
		res.Value, res.Err = m.f(url)
		m.cache[url] = res
	}
	return res.Value, res.Err
}
