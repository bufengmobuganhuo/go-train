package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	Submit(Request)
	ReadyNotifier
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// 并发的engine
func (e ConcurrentEngine) Run(seeds ...Request) {
	// 负责接收解析结果
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		// 创建worker
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		// 解析结果
		result := <-out
		for _, item := range result.Items {
			// 异步发送获取到的结果，以存储
			go func(i Item) {
				e.ItemChan <- i
			}(item)
		}
		// 将新的request放入scheduler
		for _, req := range result.Requests {
			e.Scheduler.Submit(req)
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// 通知有worker空闲
			ready.WorkerReady(in)
			// 从in channel中获取到request
			request := <-in
			parseResult, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}
