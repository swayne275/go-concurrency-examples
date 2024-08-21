package apis

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func makeAPICall(service string, ch chan<- bool) {
	time.Sleep(200 * time.Millisecond)

	fmt.Println("Making API call to", service)

	ch <- true
}

func Execute(ctx context.Context) {
	start := time.Now()
	services := []string{"service1", "service2", "service3", "service4"}

	ch := make(chan bool)
	wg := sync.WaitGroup{}

	for _, service := range services {
		wg.Add(1)
		go func() {
			defer wg.Done()
			makeAPICall(service, ch)
		}()
	}

	go func() {
		// clean up resources
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println("Result:", result)
	}

	fmt.Println("Time taken:", time.Since(start))
}
