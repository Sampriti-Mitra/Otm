package main

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00005, Down00005)
}

func Up00005(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`create table collab (
			id int not null auto_increment,
			project_title varchar(100),
			videolink varchar(100),
			created_by varchar(100),
			members json,
			members_status json,
		    created_at timestamp NULL DEFAULT NULL,
            updated_at timestamp NULL DEFAULT NULL,
            deleted_at timestamp NULL DEFAULT NULL,
		    primary key(id),
		    UNIQUE KEY project_title (project_title, created_by)
		);`)
	return err
}

func Down00005(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("drop table collab")
	return err
}
