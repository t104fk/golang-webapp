package config

import (
  "gopkg.in/alecthomas/kingpin.v2"
)

var Cfg *Config

type Config struct {
  Environment string
  DbHost string
  DbPort string
  DbName string
  DbUser string
  DbPass string
}

func Load(app *kingpin.Application) {
  Cfg = &Config{}

  app.Flag("env", "Server Environment").
    Default("development").
    OverrideDefaultFromEnvar("ENVIRONMENT").
    StringVar(&Cfg.Environment)
  app.Flag("dbhost", "Database Hostname").
    Default("localdb").
    OverrideDefaultFromEnvar("DB_HOST").
    StringVar(&Cfg.DbHost)
  app.Flag("dbport", "Database Port").
    Default("5432").
    OverrideDefaultFromEnvar("DB_PORT").
    StringVar(&Cfg.DbPort)
  app.Flag("dbname", "Database Name").
    Default("postgres").
    OverrideDefaultFromEnvar("DB_NAME").
    StringVar(&Cfg.DbName)
  app.Flag("dbuser", "Database User").
    Default("postgres").
    OverrideDefaultFromEnvar("DB_USER").
    StringVar(&Cfg.DbUser)
  app.Flag("dbpass", "Database Password").
    Default("postgres").
    OverrideDefaultFromEnvar("DB_PASS").
    StringVar(&Cfg.DbPass)
}
