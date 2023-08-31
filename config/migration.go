package config

import (
	"fmt"
	"log"
	"myadmin/app/entity"
)

func Migration(db Database) {
	fmt.Println("Process Migrating...")

	err := db.DB.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.Menu{},
		&entity.Privilege{},
		&entity.UserRole{},
		&entity.RoleMenu{},
		&entity.Permission{},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database migrated successfully...")
}
