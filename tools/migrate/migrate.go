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
	"github.com/pkg/errors"
	"github.com/romanyx/polluter"
	_ "gorm.io/driver/postgres"

	"github.com/ut080/bcs-portal/internal/config"
	"github.com/ut080/bcs-portal/internal/database"
	"github.com/ut080/bcs-portal/internal/logging"
)

func reset(pollute bool, logger logging.Logger) error {
	err := down(logger)
	if err != nil && err.Error() != migrate.ErrNoChange.Error() {
		return errors.WithStack(err)
	}
	err = up(pollute, logger)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func down(logger logging.Logger) error {
	logger.Info().Str("database", config.GetString("database.name")).Msg("migrating database down")
	m, _, err := setup(logger)
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.Down()
	if err != nil {
		return errors.WithStack(err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		return errors.WithStack(err)
	}
	logger.Info().Str("database", config.GetString("database.name")).Uint("version", version).Bool("dirty", dirty).Msg("migration complete")

	return nil
}

func up(pollute bool, logger logging.Logger) error {
	logger.Info().Str("database", config.GetString("database.name")).Msg("migrating database up")
	m, p, err := setup(logger)
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.Up()
	if err != nil {
		return errors.WithStack(err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		return errors.WithStack(err)
	}
	logger.Info().Str("database", config.GetString("database.name")).Uint("version", version).Bool("dirty", dirty).Msg("migration complete")

	if pollute {
		err = seed(p, logger)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func seed(polluter *polluter.Polluter, logger logging.Logger) error {
	if polluter == nil {
		_, p, err := setup(logger)
		if err != nil {
			return errors.WithStack(err)
		}
		polluter = p
	}

	logger.Info().Str("database", config.GetString("database.name")).Msg("seeding database")
	seedPath := config.GetString("database.migration.seed")

	f, err := os.Open(seedPath)
	if err != nil {
		logger.Error().Err(err).Str("database", config.GetString("database.name")).Msg("unable to read database seed file")
	}
	defer f.Close()

	err = polluter.Pollute(f)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func setup(logger logging.Logger) (*migrate.Migrate, *polluter.Polluter, error) {
	dbURL := database.GetDBUrl()
	logger.Debug().Str("url", dbURL).Msg("database URL generated")
	migrationsPath := config.GetString("database.migration.source")
	logger.Debug().Str("path", migrationsPath).Msg("migrations path identified")

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
	loglevel := flag.String("loglevel", "info", "set the loglevel for the tool")
	pollute := flag.Bool("seed", false, "seed database with test data after migration")
	flag.Parse()

	cmd := flag.Arg(0)
	if cmd == "" {
		fmt.Println("Usage: go run migrate.go [--seed] [--loglevel=debug|info|warn|error] [up|down|reset|seed]")
		os.Exit(1)
	}

	config.InitConfig()
	logging.InitLogging(*loglevel, true)
	logger := logging.Logger{}

	var err error
	switch strings.ToLower(cmd) {
	case "up":
		err = up(*pollute, logger)
	case "down":
		err = down(logger)
	case "reset":
		err = reset(*pollute, logger)
	case "seed":
		err = seed(nil, logger)
	default:
		fmt.Printf("Invalid command: %s\nValid commands are [up|down|reset]\n", cmd)
	}

	if err != nil {
		logger.Error().Err(err).Str("command", cmd).Msg("problem encountered while running command")
		os.Exit(1)
	}

	os.Exit(0)
}
