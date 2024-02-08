package database

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ut080/bcs-portal/internal/config"
)

func GetDBUrl() string {
	if config.GetBool("database.ssl") {
		return fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			config.GetString("database.user"),
			config.GetString("database.password"),
			config.GetString("database.host"),
			config.GetString("database.port"),
			config.GetString("database.name"),
		)
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.GetString("database.user"),
		config.GetString("database.password"),
		config.GetString("database.host"),
		config.GetString("database.port"),
		config.GetString("database.name"),
	)
}

func GetDB() (*sql.DB, error) {
	url := GetDBUrl()
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}

func GetGormDB() (*gorm.DB, error) {
	db, err := GetDB()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	gormDB, err := gorm.Open(
		postgres.New(
			postgres.Config{
				Conn: db,
			},
		),
		&gorm.Config{},
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return gormDB, nil
}
