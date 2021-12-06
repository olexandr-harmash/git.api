package app

type Config struct {
	AccessToken    string
	Own            string
	RepositoryName string
}

type Package struct {
	FullName   string
	StarsCount int
	ForksCount int
}
