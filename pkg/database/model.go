package database

type Config struct {
	URL string
}

type DataBase interface {
	Open()
	Close()

	Read()
	Write()
}
