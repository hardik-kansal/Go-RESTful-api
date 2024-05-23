package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)
type Sqlstorage struct{
	db *sql.DB
}
func connection(cfg mysql.Config)*Sqlstorage{
	db,err:=sql.Open("mysql",cfg.FormatDSN())
	if err!=nil{
		log.Fatal(err)
	}
	err=db.Ping() // verifying connection
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("connected to mysql")
	return &Sqlstorage{db:db}

}
func (s *Sqlstorage) init()(*sql.DB,error){
return s.db,nil
}
func (s *Sqlstorage) Init() (*sql.DB, error) {
	if err := s.createUsersTable(); err != nil {
		return nil, err
	}

	if err := s.createTasksTable(); err != nil {
		return nil, err
	}

	return s.db, nil
}

func (s *Sqlstorage) createUsersTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			email VARCHAR(255) NOT NULL,
			firstName VARCHAR(255) NOT NULL,
			lastName VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			UNIQUE KEY (email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	return err
}
func (s *Sqlstorage) createTasksTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			status ENUM('TODO', 'IN_PROGRESS', 'IN_TESTING', 'DONE') NOT NULL DEFAULT 'TODO',
			projectId INT UNSIGNED NOT NULL,
			AssignedToID INT UNSIGNED NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			FOREIGN KEY (AssignedToID) REFERENCES users(id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	return err
}