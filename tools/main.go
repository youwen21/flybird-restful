package main

import (
	"fmt"
	"gofly/app/service/wx_mini/wx_mini_def"
)

func main() {
	pa := wx_mini_def.Routers[2].Path()
	fmt.Println(pa)
}
