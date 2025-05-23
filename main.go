package main

import (
	"fmt"

	"github.com/liqianbro/Ginit/cmd/ginit"
)

func main() {
	if err := ginit.Execute(); err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
