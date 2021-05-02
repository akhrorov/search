package cmd

import (
	"context"
	"log"
)

func main() {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	files := []string{
		"1. aaa, bbb, ccc\n2. ddd, eee, fff\n3. ggg, hhh, iii\n",
		"4. jjj, kkk, lll\n5. xxx, nnn, ooo\n6. ppp, qqq, rrr\n",
		"7. sss, ttt, uuu\n8. vvv, www, xxx\n9. yyy, bbb, zzz\n",
	}
	result := <-search.All(ctx, "xxx", files)
	cancel()
	log.Print(result)

}
