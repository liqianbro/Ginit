package main

import (
	"Ginit/cmd/ginit"
	"fmt"
)

func main() {
	if err := ginit.Execute(); err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
