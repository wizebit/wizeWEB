package services

import (
	"path/filepath"
	"log"
	"github.com/BurntSushi/toml"
)

type dbConfig struct {
	Server string
	Port string
	Name string
	User string
	Password string
}

//read db.toml config file
func GetDbConfig() (dbconf *dbConfig) {
	configfile, err := filepath.Abs("conf/db.toml")

	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	if _, err := toml.DecodeFile(configfile, &dbconf); err != nil {
		log.Fatal(err)
	}

	return dbconf
}