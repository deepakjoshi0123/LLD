package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

// func main() {
// 	startTime := time.Now()
// 	noOfJobs := 100
// 	go allocate(noOfJobs)
// 	done := make(chan bool)
// 	go result(done)
// 	noOfWorkers := 10
// 	createWorkerPool(noOfWorkers)
// 	<-done
// 	endTime := time.Now()
// 	diff := endTime.Sub(startTime)
// 	fmt.Println("total time taken ", diff.Seconds(), "seconds")
// }

// allocate() creates jobs and puts them into the jobs channel; after all jobs are added, it closes jobs.
// Workers continuously take jobs from jobs, process them, and put the answers into the results channel.
// result() waits on results — an empty channel does not mean finished; it simply blocks and waits.
// Each worker calls wg.Done() when it has finished all available jobs.
// When all workers finish, wg.Wait() returns, and we close results.
// Closing results tells result() "no more results are coming", so its loop ends and main() continues.

// The Go runtime scheduler decides which goroutine runs on which OS thread.
// Simple analogy
// Think:
// - Process = Restaurant 🏪
// - OS Thread = Chef 👨‍🍳
// - Goroutine = Order 🍔

// Thread is managed by the OS.
// Goroutine is managed by Go.
// Go runs many goroutines on a smaller number of OS threads.

// So who does what?
// CPU Core
// "I'll execute instructions."

// OS
// "I'll schedule OS threads onto CPU cores."

// Go Runtime
// "I'll schedule goroutines onto OS threads."

// is it possible 1 cpu core can run multiple go routines at a time ?
// Yes — but not at the exact same instant on a single CPU core.

// A single CPU core can execute one instruction stream at a time. However, it can rapidly switch between goroutines, making it look like multiple goroutines are running simultaneously.

// "A CPU core executes one OS thread at a time, and that thread may be running a goroutine."
// Also, a process is a container for threads; the CPU doesn't directly execute a process.

// A CPU core directly executes OS threads, not processes or goroutines.
// When a process starts, it normally has at least one initial thread. You can call it the main/initial thread.

// COMPUTER
// │
// ┌──────┴──────┐
// │  CPU / Cores │
// └──────┬──────┘
// │
// ┌─────────────┼─────────────┐
// ▼             ▼             ▼
// Core 1         Core 2        Core 3
// ▲             ▲             ▲
// │             │             │
// OS schedules  OS schedules  OS schedules
// │             │             │
// ▼             ▼             ▼
// OS Thread 1    OS Thread 2    OS Thread 3
// ▲             ▲             ▲
// │             │             │
// │       Go Runtime Scheduler
// │             │
// └─────────────┼─────────────
// │
// ┌──────┴──────┐
// │  Go Process │
// └──────┬──────┘
// │
// ┌───────────┼───────────┐
// ▼           ▼           ▼
// Goroutine 1  Goroutine 2  Goroutine 3
// │           │           │
// └───────────┼───────────┘
// │
// Go Runtime decides
// which goroutine runs
// on which OS thread

// can we create sperate thread in go application not go ruotine ?
// Yes, technically you can, but Go does not provide a normal built-in API to directly create/manage OS threads like you do with goroutines.

// Go decides how many OS threads are needed and schedules goroutines onto them.
// But can you force a goroutine onto one OS thread?
// Yes, using:
// runtime.LockOSThread()

// Example:
// go func() {
//     runtime.LockOSThread()
//     defer runtime.UnlockOSThread()

//     // This goroutine stays on this OS thread
// }()

// Race condition
// 	- order of execution of threads deciding the final value

// program under execution is called process
// 	- heap memory
// 	- stack memory
// 	-
//    but there is some situations where we needs to do something parallelly so we need thread
