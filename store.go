package main
import (
	"database/sql"
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
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
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
