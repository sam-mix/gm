package main

import (
	"fmt"
	"gm/internal/erl"
)

func main() {
	res := erl.Call("erl-demo@127.0.0.1", "t", "h")
	fmt.Println("res = ", res)
}
