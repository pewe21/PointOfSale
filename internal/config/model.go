package config

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Tz       string
}

type Jwt struct {
	Secret string
	Exp    int
}

type Redis struct {
	Host string
	Port string
}

type Config struct {
	Server   Server
	Database Database
	Jwt      Jwt
	Redis    Redis
}
