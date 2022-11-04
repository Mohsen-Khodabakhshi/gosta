package main

import (
	"fmt"

	"github.com/Mohsen-Khodabakhshi/gosta/core"
)

func init() {
	config, err := core.InitConfig("./build/config.yaml")

	if err != nil {
		panic("Error in load config.yaml")
	}

	core.SetPublicConfig(config)
}

func main() {

	fmt.Println(core.Config)

}
