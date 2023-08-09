package manager

import (
	"api/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type InfraManager interface {
	DbConn() *sql.DB
}

type infraManager struct {
	db  *sql.DB
	cfg config.Config
}

// make database connection
func (i *infraManager) initDB() {
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.Name)
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		log.Fatalln("\nFailed connect to database\nerr: ", err)
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("App failed to run\nerr: ", err)
		}
	}()

	if err = db.Ping(); err != nil {
		log.Fatalln("\nFailed to ping database\nerr: ", err)
	}
	i.db = db
	log.Println("\nConnected to database")
}

// return db value from infraManager struct
func (i *infraManager) DbConn() *sql.DB {
	return i.db
}

// get config then share to initDB()
func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{
		cfg: config,
	}
	infra.initDB()
	return &infra
}
