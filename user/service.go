package user

import (
	"database/sql"
)

// Service the User service
type Service struct {
	Db *sql.DB
}

// FindByID find an User by Id
func (s *Service) FindByID(id int) (User, error) {
	user := User{}
	err := s.Db.QueryRow("select u.id, u.name from user u where u.id = $1", id).Scan(&user.ID, &user.Name)
	return user, err
}

// FindAll find all Users
func (s *Service) FindAll() (Users, error) {
	rows, err := s.Db.Query("select u.id, u.name from user u")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		u := User{}
		if err = rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// Create new user
func (s *Service) Create(newUser User) (User, error) {
	stmt, err := s.Db.Prepare("insert into user (name) values ($1)")
	user := User{}
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(newUser.Name)
	id, err := res.LastInsertId()
	if err != nil {
		return user, err
	}
	user.ID = int(id)
	user.Name = newUser.Name
	return user, err
}

// Update the user
func (s *Service) Update(u User) (User, error) {
	stmt, err := s.Db.Prepare("update user set name = $1 where id = $2")
	if err != nil {
		return u, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Name, u.ID)
	if err != nil {
		return u, err
	}
	return u, nil
}

// Delete the user by Id
func (s *Service) Delete(id int) error {
	stmt, err := s.Db.Prepare("delete from User where id = $1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
