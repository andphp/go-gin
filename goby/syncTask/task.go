package utils

import (
	"sync"
)

type TaskFunc func(params ...interface{})

var taskList chan *TaskExecutor //任务列表
var once sync.Once

func init() {
	chlist := getTaskList() //得到任务列表
	go func() {
		for t := range chlist {
			doTask(t)
		}
	}()
}
func doTask(t *TaskExecutor) {
	go func() {
		defer func() {
			if t.callback != nil {
				t.callback()
			}
		}()
		t.Exec()
	}()
}
func getTaskList() chan *TaskExecutor {
	once.Do(func() {
		taskList = make(chan *TaskExecutor) //初始化
	})
	return taskList
}

type TaskExecutor struct {
	f        TaskFunc
	p        []interface{} //参数
	callback func()
}

func NewTaskExecutor(f TaskFunc, p []interface{}, callback func()) *TaskExecutor {
	return &TaskExecutor{f: f, p: p, callback: callback}
}

func (t *TaskExecutor) Exec() { //执行任务
	t.f(t.p...)
}

func Task(f TaskFunc, cb func(), params ...interface{}) {
	if f == nil {
		return
	}
	go func() {
		getTaskList() <- NewTaskExecutor(f, params, cb) //增加任务队列
	}()

}
