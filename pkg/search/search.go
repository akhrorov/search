package search

import (
	"context"
	"strings"
	"sync"
)

type Result struct {
	Phrase  string
	Line    string
	LineNum int64
	ColNum  int64
}

func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	ch := make(chan []Result)
	wg := sync.WaitGroup{}
	var result []Result
	for i, file := range files {
		wg.Add(1)
		go func(ctx context.Context, file string, ch chan<- []Result, index int) {
			defer wg.Done()
			splited := strings.Split(file, "\n")
			for _, row := range splited {
				if row == "" {
					continue
				}
				if strings.Contains(row, phrase) {
					result = append(result, Result{
						Phrase:  phrase,
						Line:    row,
						LineNum: int64(index) + 1,
						ColNum:  int64(strings.Index(file, phrase)),
					})
				}
			}
		}(ctx, file, ch, i)
	}

	go func() {
		defer close(ch)
		wg.Wait()
		ch <- result
	}()

	return ch
}
