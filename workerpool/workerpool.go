package workerpool

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func sqWorker(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
	fmt.Printf("worker %d starting...\n", id)
	for job := range jobs {
		select {
		case results <- job * job:
			time.Sleep(50 * time.Millisecond) // simulate work
		case <-ctx.Done():
			return
		}
	}
}

func Execute(ctx context.Context) {
	// note: buffered channels optional
	input := []int{1, 2, 3, 4, 5}
	numWorkers := 2
	jobs := make(chan int)

	wg := sync.WaitGroup{}
	results := make(chan int)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sqWorker(ctx, i, jobs, results)
		}()
	}

	// send jobs to workers
	go func() {
		defer close(jobs)
		for _, job := range input {
			select {
			case jobs <- job:
			case <-ctx.Done():
				return
			}
		}
	}()

	// close results channel once all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// read the results
	for res := range results {
		fmt.Println(res)
	}
}
