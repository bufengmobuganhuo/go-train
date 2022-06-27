package memo1

import (
	"io/ioutil"
	"net/http"
)

// 并发不安全的版本
type Func func(string) (interface{}, error)

type Result struct {
	Value interface{}
	Err   error
}

type memo struct {
	f     Func
	cache map[string]Result
}

func New(f Func) *memo {
	return &memo{f: f, cache: make(map[string]Result)}
}

func (m *memo) Get(url string) (interface{}, error) {
	res, ok := m.cache[url]
	if !ok {
		res.Value, res.Err = m.f(url)
		m.cache[url] = res
	}
	return res.Value, res.Err
}

func HttpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}