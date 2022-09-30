package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202208040939",
		Migrate: func(tx *gorm.DB) error {
			type User struct {
				ID        uint64    `gorm:"primaryKey"`
				FirstName string    `valid:"notnull" gorm:"varchar(60)"`
				LastName  string    `valid:"notnull" gorm:"varchar(60)"`
				DDD       string    `valid:"-" gorm:"varchar(2)"`
				Phone     string    `valid:"-" gorm:"varchar(9)"`
				Birth     time.Time `valid:"-"`
				CreatedAt time.Time `valid:"-"`
				UpdatedAt time.Time `valid:"-"`
			}
			return tx.AutoMigrate(&User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("user")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
