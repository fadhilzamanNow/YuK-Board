package config

import (
	"log"
	"yukboard/models"
)

func DropAll() {
	DB.Exec(`
		DROP TABLE IF EXISTS tasks, list_invitations, list_members, todo_lists, users CASCADE;
		DROP TYPE IF EXISTS member_role, invitation_status, task_status CASCADE;
	`)
	log.Println("Dropped all tables and types")
}

func Migrate() {
	// Create ENUMs first
	DB.Exec(`
		DO $$ BEGIN CREATE TYPE member_role AS ENUM ('owner', 'member'); EXCEPTION WHEN duplicate_object THEN NULL; END $$;
		DO $$ BEGIN CREATE TYPE invitation_status AS ENUM ('pending', 'accepted', 'declined', 'cancelled'); EXCEPTION WHEN duplicate_object THEN NULL; END $$;
		DO $$ BEGIN CREATE TYPE task_status AS ENUM ('todo', 'in_progress', 'done'); EXCEPTION WHEN duplicate_object THEN NULL; END $$;
	`)

	err := DB.AutoMigrate(
		&models.User{},
		&models.TodoList{},
		&models.ListMember{},
		&models.ListInvitation{},
		&models.Task{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migration completed")
}
