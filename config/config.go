package config

type Config struct {
	Server *Server
	DB     *Database
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Dialect string
	Host    string
	Port    string
}
