package gocontext

import (
	"context"
	"fmt"
)

// transform variable
func UseContextVar(ctx context.Context) {

	id, ok := ctx.Value("id").(string)
	if ok {
		fmt.Println(id)
	} else {
		fmt.Println("error")
	}

	name, ok := ctx.Value("name").(string)
	if ok {
		fmt.Println(name)
	} else {
		fmt.Println("error")
	}

}

func TransfromVar() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "id", "linshukai")
	ctx = context.WithValue(ctx, "name", "China")
	UseContextVar(ctx)

	ctx, cancel := context.WithCancel(ctx)
	cancel()

}
