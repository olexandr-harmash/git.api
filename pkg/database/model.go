package database

type Config struct {
	URL string
}

type DataBase interface {
	Open()
	Close()

	Write()
	Read()
}
