package migration

import (
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Up(db *gorm.DB) {
	sqlDB, _ := db.DB()
	db.Dialector.Name()
	migrations := &migrate.FileMigrationSource{
		Dir: "./migration",
	}

	n, err := migrate.Exec(sqlDB, db.Dialector.Name(), migrations, migrate.Up)
	if err != nil {
		log.Errorf(fmt.Sprintf("Failed to migrate up: %s", err))
		return
	}
	fmt.Println(n)
}

func Down(db *gorm.DB) {
	sqlDB, _ := db.DB()
	db.Dialector.Name()
	migrations := &migrate.FileMigrationSource{
		Dir: "./migration",
	}

	n, err := migrate.Exec(sqlDB, db.Dialector.Name(), migrations, migrate.Down)
	if err != nil {
		log.Errorf(fmt.Sprintf("Failed to migrate down: %s", err))
		return
	}
	fmt.Println(n)
}
