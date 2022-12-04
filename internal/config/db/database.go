package db

type DatabaseList struct {
	Dbo Database
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Adapter  string
}
