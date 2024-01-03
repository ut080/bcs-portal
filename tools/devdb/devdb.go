package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/romanyx/polluter"
	_ "gorm.io/driver/postgres"
)

func reset(pollute bool) error {
	err := down()
	if err != nil && err.Error() != migrate.ErrNoChange.Error() {
		return errors.WithStack(err)
	}
	err = up(pollute)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func down() error {
	fmt.Println("Migrating dev database down...")
	m, _, err := setup()
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.Down()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func up(pollute bool) error {
	fmt.Println("Migrating dev database up...")
	m, p, err := setup()
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.Up()
	if err != nil {
		return errors.WithStack(err)
	}

	if pollute {
		err = seed(p)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func seed(polluter *polluter.Polluter) error {
	if polluter == nil {
		_, p, err := setup()
		if err != nil {
			return errors.WithStack(err)
		}
		polluter = p
	}

	fmt.Println("Seeding database...")
	seedPath := os.Getenv("DEV_SEED")

	f, err := os.Open(seedPath)
	if err != nil {
		fmt.Printf("Unable to read database seed file: %s", err)
	}
	defer f.Close()

	err = polluter.Pollute(f)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func setup() (*migrate.Migrate, *polluter.Polluter, error) {
	dbURL := os.Getenv("DEV_DATABASE_URL")
	migrationsPath := os.Getenv("DEV_MIGRATIONS")

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	p := polluter.New(polluter.PostgresEngine(db), polluter.YAMLParser)

	return m, p, nil
}

func main() {
	pollute := flag.Bool("seed", false, "seed database with test data after migration")
	flag.Parse()

	cmd := flag.Arg(0)
	if cmd == "" {
		fmt.Println("Usage: go run devdb.go [--seed] [up|down|reset|seed]")
		os.Exit(1)
	}

	if err := godotenv.Load(); err != nil {
		fmt.Println("Could not load .env, attempting to continue")
	}

	var err error
	switch strings.ToLower(cmd) {
	case "up":
		err = up(*pollute)
	case "down":
		err = down()
	case "reset":
		err = reset(*pollute)
	case "seed":
		err = seed(nil)
	default:
		fmt.Printf("Invalid command: %s\nValid commands are [up|down|reset]\n", cmd)
	}

	if err != nil {
		fmt.Printf("Problem encountered while running %s: %s", cmd, err)
		os.Exit(1)
	}

	os.Exit(0)
}
