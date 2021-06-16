package loop_task

import (
	"github.com/astaxie/beego/logs"
	"time"
)


// 循环执行的任务
type LoopTask interface {
	Run()
	Stop()
}

//
type TaskFunc func() error

type BaseLoopTask struct {
	JobChan        chan TaskFunc
	Interval       int // 间隔时间 单位秒 定时向chan中发数据
	SourceInterval int // 初始时间间隔
	f              TaskFunc
}

func NewBaseLoopTask(chanLen int, intervalTime int, f TaskFunc) *BaseLoopTask {
	task := &BaseLoopTask{
		JobChan:        make(chan TaskFunc, chanLen),
		Interval:       intervalTime,
		f:              f,
		SourceInterval: intervalTime,
	}
	go task.create()
	return task
}

func (c *BaseLoopTask) create() {
	for {
		c.JobChan <- c.f
		time.Sleep(time.Duration(c.Interval) * time.Second)

	}
}

func (c *BaseLoopTask) Stop() {

}

func (c *BaseLoopTask) Run() {
	go func() {
		for {
			select {
			case f := <-c.JobChan:
				err := f()
				if err != nil && c.Interval < 600 {
					c.Interval++
					logs.Warn("执行出错 间隔时间++", err)
				} else {
					c.Interval = c.SourceInterval
				}
			}
		}
	}()
}
