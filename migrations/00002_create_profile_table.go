package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00002, Down00002)
}

func Up00002(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`create table profile (
			id int not null auto_increment,
			videolink varchar(750),
			title varchar(200),
			tags varchar(200),
			created_by varchar(255),
		    created_at timestamp NULL DEFAULT NULL,
            updated_at timestamp NULL DEFAULT NULL,
            deleted_at timestamp NULL DEFAULT NULL,
            applauded_by json DEFAULT NULL,
  			applause int DEFAULT '0',
		    primary key(id)
		);`)
	return err
}

func Down00002(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("drop table profile")
	return err
}
