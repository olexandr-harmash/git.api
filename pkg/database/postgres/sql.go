package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/olexandr-harmash/git.api/pkg/database"
)

type Postgres struct {
	Config *database.Config
	DB     *sql.DB
}

func (p *Postgres) Open() {
	db, err := sql.Open("postgres", p.Config.URL)
	if err != nil {
		fmt.Printf("Problem in open data base %v\n", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Problem in ping data base %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("New DB \n")
	p.DB = db
}

func (p *Postgres) Close() {
	err := p.DB.Close()
	if err != nil {
		fmt.Printf("Problem in close data base %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Close DB \n")
}

func (p *Postgres) Write(name string, stars, forks int) sql.Result {
	stmt := fmt.Sprintf("INSERT INTO repo(full_name, stars_count, forks_count) VALUES('%s', '%d', '%d');", name, stars, forks)

	result, err := p.DB.Exec(stmt)

	if err != nil {
		fmt.Printf("Problem in exec data base %v\n", err)
		os.Exit(1)
	}

	return result
}

func (p *Postgres) Read() {
}
