package gocontext

import (
	"context"
	"fmt"
	"sync"
)

func handleZhangsan(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx = context.WithValue(ctx, "name1", "zhangsan")
	fmt.Println("name1: ", ctx.Value("name1"))
}

func handleLining(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx = context.WithValue(ctx, "name1", "linsss")
	fmt.Println("name2: ", ctx.Value("name1"))
}

// 测试context的是否线程安全
// conetxt线程安全，因为每个使用withvalue的相当于套了一层Context
// 获取值的时候递归沿着context寻找value
func HandlerUserValue() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "linshukai")
	ctx.Done()

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		handleZhangsan(ctx, wg)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		handleLining(ctx, wg)
	}

	wg.Wait()
}
