package main

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

var ErrInvalidMParam = errors.New("invalid params. Give chance to mistake for app :)")

type Task func() error

func Run(tasks []Task, n, m int) (e error) {
	if m <= 0 {
		return ErrInvalidMParam
	}

	var count int32

	taskCh := make(chan Task)
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for {
				t, ok := <-taskCh
				if !ok {
					return
				}
				if err := t(); err != nil {
					atomic.AddInt32(&count, 1)
				}
			}
		}()
	}

	wg.Add(n)
	for _, task := range tasks {
		if int32(m) > 0 && atomic.LoadInt32(&count) >= int32(m) {
			e = ErrErrorsLimitExceeded
			break
		}
		taskCh <- task
		continue
	}
	close(taskCh)
	wg.Wait()

	return
}
