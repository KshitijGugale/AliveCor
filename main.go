package main

import (
	"fmt"
	"time"
)

type Task struct {
	Id int
	IsCompleted bool
	Status string // untouched, completed, failed, timeout
	CreationTime time.Time // when was the task created
	TaskData string // field containing data about the task
}

func main() {
	var t [10]Task

	// Add task
	for i:= 0; i < 10; i++ {
		t[i].Id = i
		t[i].IsCompleted = false
		t[i].Status = "untouched"
		t[i].CreationTime = time.Now()
		t[i].TaskData = "Added  to the queue"
	}

	ch := make(chan Task, 10)
	for i, _ := range t {
		ch <- t[i]
	}
	close(ch)

	// Execute tasks
	go Executer(ch)
	time.Sleep(500 * time.Millisecond)

}

func Executer(ch chan Task) {
	fmt.Println("Executing all the tasks...")

	for _ = range ch {
		v := <- ch
		fmt.Println(v)
		clean(v)
	}

}

func clean(v Task) {
	v.IsCompleted = true
	v.Status = "completed"
	v.TaskData = "Removed from queue"

	fmt.Println(v)
}