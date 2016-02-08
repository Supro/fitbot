package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type database struct {
	Connection *sql.DB
}

var db_instance *database
var db_once sync.Once

func GetDatabaseInstance() *database {
	db_once.Do(func() {
		cfg := GetConfigurationInstance()

		var cfgstr string

		if cfg.Host != "" {
			cfgstr += "host=" + cfg.Host + " "
		}

		if cfg.DBName != "" {
			cfgstr += "dbname=" + cfg.DBName + " "
		}

		if cfg.User != "" {
			cfgstr += "user=" + cfg.User + " "
		}

		if cfg.Password != "" {
			cfgstr += "password=" + cfg.Password + " "
		}

		if cfg.SSLMode != "" {
			cfgstr += "sslmode=" + cfg.SSLMode + " "
		}

		db, err := sql.Open("postgres", cfgstr)
		if err != nil {
			log.Fatal(err)
		}

		db_instance = &database{Connection: db}
	})

	return db_instance
}
