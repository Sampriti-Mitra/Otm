package main

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00004, Down00004)
}

func Up00004(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`create table about (
			id int not null auto_increment,
			about varchar(750),
			user_id varchar(200),
			followers int,
			following int,
		    created_at timestamp NULL DEFAULT NULL,
            updated_at timestamp NULL DEFAULT NULL,
            deleted_at timestamp NULL DEFAULT NULL,
		    primary key(id),
		    UNIQUE KEY user_id (user_id)
		);`)
	return err
}

func Down00004(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("drop table about")
	return err
}
