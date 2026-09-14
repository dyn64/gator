package config

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

var defConf = Config{
	DbURL:           "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
	CurrentUserName: "secret",
}
