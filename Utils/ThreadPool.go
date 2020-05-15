package Utils

import (
	"time"
)

type task func(curTime time.Time)

type Worker struct {
	// 这里线程池的二级channel可以这么理解，线程池中多个channel是用来存放woker线程，用来控制线程的数量，而每个线程的结构体又是一个channel，这个channel的类型是task，用来等待任务的发生
	workerPool  chan chan task // 线程池，即woker所属的线程池
	taskChannel chan task      // 任务通道
	quit        chan bool      // 退出通道
}

// 新建一个woker线程
func newWorker(workPool chan chan task) *Worker {
	return &Worker{
		workerPool:  workPool, // 表示work所在的线程池
		taskChannel: make(chan task),
		quit:        make(chan bool),
	}
}

// 给线程定义一个start方法，表示监听到有任务来了开始干活
func (this *Worker) start() {
	go func() {
		// 表示线程池中的某一个woker开始处理任务了，这里线程池如果满了就不再接收新任务了，会在这里阻塞
		// 这里就是一级channel 用来限制线程池中worker线程的使用
		this.workerPool <- this.taskChannel
		// 二级channel，这里开始处理任务，channel中没有任务就阻塞,直到该线程被停止
		select {
		case taskObj := <-this.taskChannel:
			taskObj(time.Now())
		case <-this.quit:
			return
		}
	}()
}

// 线程停止工作
func (this *Worker) stop() {
	this.quit <- true
}

type Dispatcher struct {
	workPool  chan chan task
	maxNum    int       // 线程池中线程数的最大数量
	taskQueue chan task // 任务通道
}

// 新建一个任务分发器
func NewDispatcher(maxWorkerNum int) *Dispatcher {
	return &Dispatcher{
		workPool:  make(chan chan task, maxWorkerNum),
		maxNum:    maxWorkerNum,
		taskQueue: make(chan task),
	}
}

// 添加任务
func (this *Dispatcher) addTask(t task) {
	this.taskQueue <- t
}

// 分配任务
func (this *Dispatcher) dispatcher() {
	for {
		select {
		// 从任务队列中取出一个任务
		case taskObj := <-this.taskQueue:
			go func(t task) {
				// 从线程池中取出一个线程
				workerChannel := <-this.workPool
				workerChannel <- taskObj
			}(taskObj)
		}
	}
}

// 启动任务分配器，任务分配球开始运行并分发任务
func (this *Dispatcher) Run() {
	// 新建工作线程
	for i := 0; i < int(this.maxNum); i++ {
		workerObj := newWorker(this.workPool)
		// 启动线程
		workerObj.start()
	}
	// 分发任务
	go this.dispatcher()
}
