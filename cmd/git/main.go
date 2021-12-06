package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/olexandr-harmash/git.api/internal/app"
	"github.com/olexandr-harmash/git.api/internal/app/gitapi"
	"github.com/olexandr-harmash/git.api/pkg/database"
	"github.com/olexandr-harmash/git.api/pkg/database/postgres"
)

const (
	gitConfig string = "./configs/git_config.toml"
	dbConfig  string = "./configs/db_config.toml"
)

type App struct {
	AppConfig *app.Config
	DBConfig  *database.Config
}

func main() {

	app := &App{
		&app.Config{},
		&database.Config{},
	}

	DecodeConfig(gitConfig, app.AppConfig)
	DecodeConfig(dbConfig, app.DBConfig)

	db := &postgres.Postgres{
		Config: app.DBConfig,
	}

	gAPI := gitapi.New(app.AppConfig, db)
	gAPI.DB.Open()
	defer gAPI.DB.Close()

	ctx := gAPI.Auth()

	pack := gAPI.Get(ctx)

	gAPI.DB.Write(pack.FullName, pack.StarsCount, pack.ForksCount)
}

func DecodeConfig(file string, config interface{}) {
	_, err := toml.DecodeFile(file, config)

	if err != nil {
		fmt.Printf("Problem in getting config information %v\n", err)
		os.Exit(1)
	}
}
