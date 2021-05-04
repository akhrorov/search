package main

import (
	"context"
	"github.com/akhrorov/search/pkg/search"
	"log"
)

func main() {
	root := context.Background()
	files := []string{"data/first.txt", "data/second.txt", "data/third.txt"}
	//resultAll := <-search.All(root, "aaa", files)
	//for _, val := range resultAll{
	//	log.Print(val)
	//}
	resultAny := <-search.Any(root, "aaa", files)
	log.Print(resultAny)
}
