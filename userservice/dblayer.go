package userservice

import (
	"fmt"

	"scratch/mockExample/util"

	"github.com/go-pg/pg"
)

// Storage ...
type Storage interface{
	CreateUpdate(u util.User, db pg.DB) (string, error)
	Get(u *util.User, db pg.DB) (string, error)
	Delete(u *util.User, db pg.DB) (string, error)
}

// Store ...
type Store struct{}

// CreateUpdate ...
func (s *Store) CreateUpdate(u util.User, db pg.DB) (string, error) {

	var msg string

	_, err := db.Model(&u).
		OnConflict("(name) DO UPDATE").
		Set("email = ?", u.Email).
		Insert()
	if err != nil {
		fmt.Println("error inserting or updating db", err)
		msg = err.Error()
		return msg, err
	}

	msg = fmt.Sprintf("...successfully inserted or updated %v %v", u.Name, u.Email)

	return msg, nil
}

// Get ...
func (s *Store) Get(u *util.User, db pg.DB) (string, error) {

	var msg string

	err := db.Model(u).Where("Name = ?", u.Name).Select()

	if err != nil {
		fmt.Println("error with get db", err)
		msg = err.Error()
		return msg, err
	}

	msg = fmt.Sprintf("%+v", *u)

	return msg, nil
}

// Delete ...
func (s *Store) Delete(u *util.User, db pg.DB) (string, error) {

	var msg string

	res, err := db.Model(u).Where("Name = ?", u.Name).Delete()

	if err != nil {
		fmt.Println("error with get db", err)
		msg = err.Error()
		return msg, err
	}

	affected := res.RowsAffected()

	if affected == 0 {
		fmt.Println("did not find record to delete")
		msg = "did not find record to deleted"
		return msg, nil
	}

	msg = fmt.Sprintf("successfully deleted %v", u.Name)

	return msg, nil
}
