package runner

import (
	"os"
	"time"
	"errors"
	"os/signal"
)

type Runner struct {
	interrupt chan os.Signal // 通道报告从操作系统发送的信号
	complete chan error // 通道报告处理任务已经完成
	timeout <-chan time.Time // timeout 报告处理任务已经超时 只能接受
	tasks []func(int) //持有一组以索引顺序依次执行的函数
}

var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")

// 返回一个新的准备使用的Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete: make(chan error),
		timeout: time.After(d),
	}
}

// Add 将一个任务附加到Runner上。
func (r *Runner)Add(tasks ...func(int))  {
	r.tasks = append(r.tasks, tasks...)
}

// Start执行所有任务，并监视通道事件
func (r *Runner)Start() error {
	// 我们希望接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)
	// 用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

// run 执行每一个已注册的任务
func (r *Runner)run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt(){
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

// 验证是否接收到了中断信号
func (r *Runner)gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	// 继续正常运行
	default:
		return false
	}
}