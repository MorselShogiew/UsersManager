package repos

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/MorselShogiew/UsersManager/config"
	"github.com/avast/retry-go"
	_ "github.com/lib/pq"
)

type DBManager struct {
	Ok bool
	db *sql.DB
}

func (m *DBManager) Connect(conf *config.PostgresConfig) {
	var conninfo string = fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		conf.Host, conf.Port, conf.DB, conf.Username, conf.Password)

	var err error
	defer func() {
		m.Ok = err == nil
	}()

	err = retry.Do(
		func() error {
			log.Println("Trying to connect to the DB...")
			db, err := sql.Open("postgres", conninfo)
			if err != nil {
				return err
			}

			err = db.Ping()
			if err != nil {
				db.Close()
				return err
			}

			m.db = db
			return nil
		},
		retry.Attempts(5),
		retry.Delay(time.Second),
		retry.DelayType(retry.BackOffDelay),
	)
}

func (m *DBManager) Close() {
	m.db.Close()
}

func connectDB(conninfo string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
