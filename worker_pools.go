package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 0; i < numJobs; i++ {
		<-results
	}
}
