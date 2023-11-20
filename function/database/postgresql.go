package database

var connection string

func init() {
	connection = "PostgreSQL"
}

func GetDatabase() string {
	return connection
}
