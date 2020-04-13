package main

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00003, Down00003)
}

func Up00003(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`create table follower_request (
			id int not null auto_increment,
			request_by varchar(100),
			request_to varchar(100),
			status varchar(20) default 'pending',
		    created_at timestamp NULL DEFAULT NULL,
            updated_at timestamp NULL DEFAULT NULL,
            deleted_at timestamp NULL DEFAULT NULL,
		    primary key(id),
		    UNIQUE KEY request_status (request_by,request_to,status)
		);`)
	return err
}

func Down00003(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("drop table follower_request")
	return err
}
