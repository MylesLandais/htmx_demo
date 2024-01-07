package config

import "os"

type Config struct {
  Db string 
  Username string
  Password string
  Host string
}

func GetConfig() Config {
  config := Config{}
  config.Db = os.Getenv("PG_DATABASE")
  config.Username = os.Getenv("PG_USERNAME")
  config.Password = os.Getenv("PG_PASSWORD")
  config.Host = os.Getenv("PG_HOST")
  return config
}
