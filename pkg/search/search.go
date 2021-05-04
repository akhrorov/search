package search

import (
	"context"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

type Result struct {
	Phrase  string
	Line    string
	LineNum int64
	ColNum  int64
}

func Any(root context.Context, phrase string, files []string) <-chan Result {
	ch := make(chan Result)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(root)
	for i, filePath := range files {
		wg.Add(1)
		go func(ctx context.Context, filePath string, ch chan Result, index int) {
			defer wg.Done()
			result := FindAnyMatchesTextInFile(phrase, filePath)
			select {
			case <-ctx.Done():
				return
			default:
				if result != nil {
					cancel()
					ch <- *result
				}
			}
		}(ctx, filePath, ch, i)
	}

	go func() {
		wg.Wait()
		cancel()
		close(ch)
	}()

	return ch
}

func FindAnyMatchesTextInFile(phrase, fileName string) (res *Result) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("error", err)
		return
	}

	file := string(data)

	splited := strings.Split(file, "\n")
	for i, line := range splited {
		if strings.Contains(line, phrase) {
			res = &Result{
				Phrase:  phrase,
				Line:    line,
				LineNum: int64(i + 1),
				ColNum:  int64(strings.Index(line, phrase)) + 1,
			}
			return
		}
	}
	return
}

func All(root context.Context, phrase string, files []string) <-chan []Result {
	ch := make(chan []Result)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(root)
	for i, filePath := range files {
		wg.Add(1)
		go func(ctx context.Context, filePath string, ch chan<- []Result, index int) {
			defer wg.Done()
			result := FindAllMatchesTextInFile(phrase, filePath)
			if len(result) > 0 {
				ch <- result
			}
		}(ctx, filePath, ch, i)
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	cancel()
	return ch
}

func FindAllMatchesTextInFile(phrase, fileName string) (res []Result) {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("error", err)
		return res
	}

	file := string(data)

	splited := strings.Split(file, "\n")

	for i, line := range splited {
		if strings.Contains(line, phrase) {

			r := Result{
				Phrase:  phrase,
				Line:    line,
				LineNum: int64(i + 1),
				ColNum:  int64(strings.Index(line, phrase)) + 1,
			}

			res = append(res, r)
		}
	}

	return res
}
