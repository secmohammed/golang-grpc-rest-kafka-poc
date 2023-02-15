package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
	"gorm.io/gorm"
)

func synchronize(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "initial",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(
					&entities.Company{},
					&entities.User{},
				); err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(
					&entities.Company{},
					&entities.User{},
				)
			},
		},
	})
	return m.Migrate()
}
