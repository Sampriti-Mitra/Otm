package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00001, Down00001)
}

func Up00001(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`create table users (
			id int not null auto_increment,
			username varchar(30),
			email varchar(100),
			name varchar(100),
			password varchar(100),
			created_by varchar(255),
		    created_at timestamp NULL DEFAULT NULL,
            updated_at timestamp NULL DEFAULT NULL,
            deleted_at timestamp NULL DEFAULT NULL,
		    primary key(id),
		    UNIQUE KEY public_id (username),
  			UNIQUE KEY email (email)
		);`)
	return err
}

func Down00001(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("drop table users")
	return err
}
