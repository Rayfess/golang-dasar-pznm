package database

var connection string

func init() {
	connection = "mySQL"
}

func GetDB() string {
	return connection
}