package migration

import (
	"crud/config"
	"fmt"

	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"

	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
)

const dbMigrationsPath = "file://migration/queries"

var migrationRunner *migrate.Migrate

func Init() {
	connectionURL := config.DbConfig().ConnectionURL()
	var err error
	migrationRunner, err = migrate.New(dbMigrationsPath, connectionURL)

	if err != nil {
		fmt.Println(err)
	}
}

func RunDatabaseMigrations() error {
	err := migrationRunner.Up()
	if err != nil {
		return err
	}
	fmt.Printf("Migration successful")
	return nil
}

func RollbackLatestMigration() error {
	err := migrationRunner.Steps(-1)
	if err != nil {
		return err
	}
	fmt.Printf("Down migrations successful")
	return nil
}
