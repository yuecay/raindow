package main

import (
	"fmt"
	"rainbow/internal/middleware/sensitive"
	_ "rainbow/internal/packed"
	// "github.com/gogf/gf/v2/os/gctx"
	// "rainbow/internal/cmd"
)

func main() {
	// cmd.Main.Run(gctx.New())
	start := sensitive.SensitiveNode{}
	nodes := start.GetInstance()
	fmt.Println(nodes[0])
}
