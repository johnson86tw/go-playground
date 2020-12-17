package main

import (
	"context"
	"fmt"
)

var cancelBefore = true

func main() {
	c, cCancel := context.WithCancel(context.Background())

	c1, c1Cancel := context.WithCancel(c)
	defer c1Cancel()

	c2, c2Cancel := context.WithCancel(c)
	defer c2Cancel()

	c11, c11Cancel := context.WithCancel(c1)
	defer c11Cancel()

	c12, c12Cancel := context.WithCancel(c1)
	defer c12Cancel()

	if cancelBefore {
		fmt.Println("start cancel c1")
		c1Cancel()
	}

	ctxs := map[string]context.Context{"c1": c1, "c11": c11, "c12": c12, "c2": c2}

	for key, ctx := range ctxs {
		var s string

		if ctx.Err() != nil {
			s = "cancelled correctly"
		} else {
			s = "not cancelled"
		}

		fmt.Println(key, " is ", s)
	}

	if !cancelBefore {
		fmt.Println("start cancel c")
		cCancel()
	}

}
