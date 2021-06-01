package main

import "time"

const (
	lowPriority  = 0
	midPriority  = 1
	highPriority = 2
)

type task struct {
	title       string
	description string
	done        bool
	category    string
	priority    int
	due         *time.Time
	completion  float64
}

type taskList struct {
	tasks []*task
}

// remaining 将返回剩余的未完成的任务列表
func (l *taskList) remaining() []*task {
	var items []*task
	for _, task := range l.tasks {
		if !task.done {
			items = append(items, task)
		}
	}
	return items
}

// done 将返回已完成的任务列表
func (l *taskList) done() []*task {
	var items []*task
	for _, task := range l.tasks {
		if task.done {
			items = append(items, task)
		}
	}
	return items
}

// dummyData
func dummyData() *taskList {
	list := &taskList{
		tasks: []*task{
			{
				title:       "Nearly done",
				description: "you can tick my checkbox and I will be marked as done and disappear",
			},
			{
				title:       "Functions",
				description: "Tap the plus icon above to add a new task, or tap the minus icon to remove this one",
			},
		},
	}
	return list
}

// add 添加新的任务,并置于任务列表顶部
func (l *taskList) add(t *task) {
	l.tasks = append([]*task{t}, l.tasks...)
}

//
