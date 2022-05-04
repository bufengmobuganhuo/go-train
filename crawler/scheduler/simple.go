package scheduler

import "mengyu.com/gotrain/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(req engine.Request) {
	// 将请求放入channel
	// 避免死循环，详见幕布
	go func() {
		s.workerChan <- req
	}()
}
