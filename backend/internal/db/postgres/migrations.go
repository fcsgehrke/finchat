package postgres

import (
	"log"

	"github.com/fcsgehrke/finchat/internal/db/entities"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func (r *PostgresRepository) RunMigrations() {
	migrator := gormigrate.New(r.db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		MigrationCreateUsersTable(),
		MigrationCreateRoomsTable(),
		MigrationCreateMessageTable(),
	})

	if err := migrator.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Printf("Migration did run successfully")
}

func MigrationCreateUsersTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "CreateUsersTable",
		Migrate: func(d *gorm.DB) error {
			return d.AutoMigrate(&entities.User{})
		},
		Rollback: func(d *gorm.DB) error {
			return d.Migrator().DropTable(&entities.User{})
		},
	}
}

func MigrationCreateRoomsTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "CreateRoomsTable",
		Migrate: func(d *gorm.DB) error {
			return d.AutoMigrate(&entities.Room{})
		},
		Rollback: func(d *gorm.DB) error {
			return d.Migrator().DropTable(&entities.Room{})
		},
	}
}

func MigrationCreateMessageTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "CreateMessageTable",
		Migrate: func(d *gorm.DB) error {
			return d.AutoMigrate(&entities.Message{})
		},
		Rollback: func(d *gorm.DB) error {
			return d.Migrator().DropTable(&entities.Message{})
		},
	}
}
