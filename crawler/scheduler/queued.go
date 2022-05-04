package scheduler

import "mengyu.com/gotrain/crawler/engine"

type QueuedScheduler struct {
	// 保存request的队列
	requestChan chan engine.Request
	// 一个worker要执行的是一个request的队列，所以这里用于存储worker
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request  {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(req engine.Request) {
	s.requestChan <- req
}

// 有一个worker ready了，将这个worker送进去
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQue []engine.Request
		var workerQue []chan engine.Request
		for {
			var activeReq engine.Request
			var activeWorker chan engine.Request
			// 如果有request，并且有空闲worker
			if len(requestQue) > 0 && len(workerQue) > 0 {
				activeWorker = workerQue[0]
				activeReq = requestQue[0]
			}
			select {
				// scheduler的request队列中有新request
			case r := <-s.requestChan:
				// 如果request队列中有新request，则放入队列
				requestQue = append(requestQue, r)
				// scheduler的worker队列中有新空闲worker
			case w := <-s.workerChan:
				// 如果有新空闲的worker，那么放入队列
				workerQue = append(workerQue, w)
				// 前两个发送成功后，这个才会执行到，才会将request发送给一个worker，
				// 上面的if会从request中取出第一个元素，那么就要移除第一个元素，以便获取下一个request
				// 取出request给worker，worker从channel中取出request去fetch和parse
			case activeWorker <- activeReq:
				workerQue = workerQue[1:]
				requestQue = requestQue[1:]
			}
		}
	}()
}
