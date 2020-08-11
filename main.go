package main

import (
	"flag"
	"fmt"
	"github.com/mirrorsge/ginGen/initialize"
	"github.com/mirrorsge/ginGen/update"
)

func main() {
	//捕捉底层错误
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ERROR: there is something wrong when processing, program exited,", err)
		}
	}()
	//解析参数
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("ERROR: the number of args is wrong,only one args is needed")
		return
	}
	mode := flag.Arg(0)
	//选择对应的处理模式
	switch mode {
	case "I":
		initialize.Init()
	case "U":
		update.Update()
	default:
		fmt.Println("ERROR: the args is not one of 'I' or 'U'")
	}
	return
}
