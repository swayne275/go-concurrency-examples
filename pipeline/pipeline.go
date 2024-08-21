package pipeline

import (
	"context"
	"fmt"
)

func intSeq(ctx context.Context, in ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)

		for _, v := range in {
			select {
			case out <- v:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func sq(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)

		for n := range in {
			select {
			case out <- n * n:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func Execute(ctx context.Context) {
	input := []int{1, 2, 3, 4, 5}

	inputCh := intSeq(ctx, input...)
	squaredCh := sq(ctx, inputCh)

	for result := range squaredCh {
		fmt.Println(result)
	}
}
