package main
import (
	"database/sql"
	"log"
)

type Store struct{
	db *sql.DB
}
func Newstore(db *sql.DB)*Store{
	return &Store{
		db:db,
	}
}
func (s* Store) CreateTask(t *Task)(*Task,error){
	rows, err := s.db.Exec("INSERT INTO tasks (name,projectId,AssignedToID) VALUES (?, ?, ?)", t.Name, t.ProjectID, t.AssignedToID)

	if err != nil {
		log.Println("eror is here 111")
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		log.Println("eror is here 222")
		return nil, err
	}

	t.ID = id
	return t, nil
}
func (s *Store) GetUserByID(id string) (*User, error) {
	var u User
	err := s.db.QueryRow("SELECT id, email, firstName, lastName, createdAt FROM users WHERE id = ?", id).Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.CreatedAt)
	return &u, err
}
func (s *Store) CreateUser(u *User) (*User, error) {
	rows, err := s.db.Exec("INSERT INTO users (email, firstName, lastName, password) VALUES (?, ?, ?, ?)", u.Email, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.ID = id
	return u, nil
}
/* id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			status ENUM('TODO', 'IN_PROGRESS', 'IN_TESTING', 'DONE') NOT NULL DEFAULT 'TODO',
			projectId INT UNSIGNED NOT NULL,
			AssignedToID INT UNSIGNED NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			FOREIGN KEY (AssignedToID) REFERENCES users(id),
			FOREIGN KEY (projectId) REFERENCES projects(id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
 */