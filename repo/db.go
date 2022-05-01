package repo

import (
	"database/sql"
	"fmt"
	"github.com/PB-Digital/ms-retail-products-info/properties"
	"github.com/go-pg/pg"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"time"
)

var Db *pg.DB

func InitDb() {
	Db = pg.Connect(&pg.Options{
		Addr:        properties.Props.DbHost + ":" + properties.Props.DbPort,
		Database:    properties.Props.DbName,
		User:        properties.Props.DbUser,
		Password:    properties.Props.DbPass,
		PoolSize:    5,
		DialTimeout: 1 * time.Minute,
		MaxRetries:  2,
		MaxConnAge:  15 * time.Minute,
	})
}

func MigrateDb() error {
	log.Info("MigrateDb.start")

	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", properties.Props.DbName,
		properties.Props.DbUser, properties.Props.DbPass, properties.Props.DbHost, properties.Props.DbPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	log.Info("Applied ", n, " migrations")
	log.Info("MigrateDb.end")
	return nil
}
