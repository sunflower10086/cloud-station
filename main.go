package main

import (
	"fmt"

	"github.com/sunflower10086/cloud-station/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
