package postgres_repo

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/romanyx/polluter"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RepositorySuite struct {
	suite.Suite
	db              *gorm.DB
	rootDatabaseURL string
	testDatabaseURL string
	databaseName    string
	migrationsURL   string
	migrator        *migrate.Migrate
	polluter        *polluter.Polluter
	seedData        string
}

func (suite *RepositorySuite) SetupSuite() {
	// Fetch URLs from environment
	databaseURL := os.Getenv("TEST_DATABASE_URL")
	suite.databaseName = os.Getenv("TEST_DATABASE_NAME")
	suite.migrationsURL = os.Getenv("TEST_MIGRATIONS")
	seedURL := os.Getenv("TEST_SEED")

	// Initialize connection to Postgres
	suite.testDatabaseURL = strings.Replace(databaseURL, "<DATABASE>", fmt.Sprintf("/%s", suite.databaseName), 1)
	db, err := sql.Open("pgx", suite.testDatabaseURL)
	if err != nil {
		panic(err)
	}

	// Initialize ORM
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	suite.db = gormDB

	// Initialize the migrator, and ensure database is migrated
	m, err := migrate.New(suite.migrationsURL, suite.testDatabaseURL)
	if err != nil {
		panic(err)
	}
	suite.migrator = m

	err = suite.migrator.Up()
	if err != nil {
		panic(err)
	}

	// Initialize seeder
	suite.polluter = polluter.New(polluter.PostgresEngine(db), polluter.YAMLParser)

	sd, err := os.ReadFile(seedURL)
	if err != nil {
		panic(err)
	}

	suite.seedData = string(sd)
}

func (suite *RepositorySuite) SetupTest() {
	db, err := suite.db.DB()
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("DELETE FROM element_members;")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("DELETE FROM elements;")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("DELETE FROM flights;")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("DELETE FROM duty_assignments;")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("DELETE FROM members;")
	if err != nil {
		panic(err)
	}

	// Re-seed database
	err = suite.polluter.Pollute(strings.NewReader(suite.seedData))
	if err != nil {
		panic(err)
	}
}

func (suite *RepositorySuite) TearDownSuite() {
	err := suite.migrator.Down()
	if err != nil {
		panic(err)
	}
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}
