package main

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/cmd/app"
	"os"
)

func main() {
	if err := app.Execute(); err != nil {
		fmt.Println("Failed: ", err.Error())
		os.Exit(1)
	}
}
