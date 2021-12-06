package app

type Config struct {
	AccessToken    string
	RepositoryName string
}

type Package struct {
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}
