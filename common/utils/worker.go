package utils

import "sync"

// RunParallelFunc should be checked with lens
func RunParallelFunc(tasks ...func() error) []error {
	var wg sync.WaitGroup
	errs := make([]error, len(tasks))

	for i, task := range tasks {
		wg.Add(1)
		go func(idx int, t func() error) {
			defer wg.Done()
			errs[idx] = t()
		}(i, task)

	}
	wg.Wait()
	return errs
}
