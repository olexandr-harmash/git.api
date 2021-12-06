package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/olexandr-harmash/git.api/internal/app"
)

func main() {
	conf := &app.Config{}

	DecodeConfig(conf)

	fmt.Println(conf.AccessToken)
	fmt.Println(conf.RepositoryName)

	//gAPI := gitapi.New(conf)

	//ctx := gAPI.Auth()

	//gAPI.Get(ctx)
}

func DecodeConfig(config *app.Config) {
	_, err := toml.DecodeFile("./config/config.toml", config)

	if err != nil {
		fmt.Printf("Problem in getting config information %v\n", err)
		os.Exit(1)
	}
}
