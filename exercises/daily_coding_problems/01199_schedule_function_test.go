package dailycodingproblems

import (
	"log"
	"testing"
	"time"
)

/**
* Implement a job scheduler which takes in a function f and an integer n,
* and calls f after n milliseconds.
* */

type Job struct {
	function func()
	timer    time.Timer
}

type Scheduler struct {
	jobs chan Job
	quit chan struct{}
}

func (scheduler *Scheduler) schedule(function func(), delay int) {
	duration := time.Duration(delay) * time.Millisecond
	job := Job{
		function: function,
		timer:    *time.AfterFunc(duration, func() { function() }),
	}
	scheduler.jobs <- job
}

func scheduleMe() {
	log.Println("Thank you for scheduling me.")
}

func (scheduler *Scheduler) run() {
	for {
		select {
		case job := <-scheduler.jobs:
			job.function()
		case <-scheduler.quit:
			return
		}
	}
}

func TestScheduler(t *testing.T) {
	scheduler := Scheduler{
		jobs: make(chan Job),
		quit: make(chan struct{}),
	}
	go scheduler.run()

	scheduler.schedule(scheduleMe, 1)
	t.Log("Function executed!")
}
