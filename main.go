package main

import (
	"fmt"
	"gm/internal/conf"
	"gm/internal/erl"
)

func main() {
	println(conf.NodeCookie)
	erlSrv, err := erl.NewServer()
	if err != nil {
		panic(fmt.Errorf("分裂进程失败: %s", err))
	}
	result, err := erlSrv.Call("ml_dev_1@127.0.0.1", "dev", "u")
	if err != nil {
		panic(fmt.Errorf("远程过程调用失败: %s", err))
	}
	fmt.Println(result)

	erlSrv.Node.Wait()
}
