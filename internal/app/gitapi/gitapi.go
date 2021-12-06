package gitapi

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/olexandr-harmash/git.api/internal/app"
	"github.com/olexandr-harmash/git.api/pkg/database/postgres"
	"golang.org/x/oauth2"
)

type GitApi struct {
	Config *app.Config
	Client *github.Client
	DB     *postgres.Postgres
}

func New(config *app.Config, db *postgres.Postgres) *GitApi {
	return &GitApi{
		Config: config,
		DB:     db,
	}
}

func (g *GitApi) Auth() context.Context {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.Config.AccessToken},
	)

	tc := oauth2.NewClient(ctx, ts)

	g.Client = github.NewClient(tc)

	return ctx
}

func (g *GitApi) Get(ctx context.Context) *app.Package {
	repo, _, err := g.Client.Repositories.Get(ctx, g.Config.Own, g.Config.RepositoryName)

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	pack := &app.Package{
		FullName:   *repo.FullName,
		ForksCount: *repo.ForksCount,
		StarsCount: *repo.StargazersCount,
	}

	return pack
}
