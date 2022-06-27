package memo3

import (
	"sync"

	"mengyu.com/gotrain/lang/goroutine/memo1"
)

type entry struct {
	res memo1.Result
	// 负责通知其他goroutine，当前entry已经创建完毕
	ready chan struct{}
}

type Memo struct {
	f     memo1.Func
	cache map[string]*entry
	mu    sync.Mutex
}

func New(f memo1.Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (m *Memo) Get(url string) (interface{}, error) {
	// 首先获取锁
	m.mu.Lock()
	e := m.cache[url]
	// 如果存储函数结果的entry不存在
	if e == nil {
		// 创建一个空对象,并放入缓存
		e = &entry{ready: make(chan struct{})}
		m.cache[url] = e
		// 因为对于map和entry的操作已经完成，所以这里可以释放锁
		m.mu.Unlock()

		// 执行真正的函数
		e.res.Value, e.res.Err = m.f(url)

		close(e.ready)
		// 如果缓存中有值（这个值可能是一个空值）
	} else {
		m.mu.Unlock()

		// 这里会一直阻塞等待，上面38行函数调用完成并赋值，所以43行可以提前释放锁
		<-e.ready
	}
	return e.res.Value, e.res.Err
}
