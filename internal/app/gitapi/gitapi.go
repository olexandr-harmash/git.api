package gitapi

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/olexandr-harmash/git.api/internal/app"
	"golang.org/x/oauth2"
)

type GitApi struct {
	config *app.Config
	client *github.Client
}

func New(config *app.Config) *GitApi {
	return &GitApi{
		config: config,
	}
}

func (g *GitApi) Auth() context.Context {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.config.AccessToken},
	)

	tc := oauth2.NewClient(ctx, ts)

	g.client = github.NewClient(tc)

	return ctx
}

func (g *GitApi) Get(ctx context.Context) {
	repo, _, err := g.client.Repositories.Get(ctx, g.config.RepositoryName, "Lessons")

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	pack := &app.Package{
		FullName:    *repo.FullName,
		Description: *repo.Description,
		ForksCount:  *repo.ForksCount,
		StarsCount:  *repo.StargazersCount,
	}

	fmt.Printf("%+v\n", pack)
}
