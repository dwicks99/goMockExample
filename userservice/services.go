package userservice

import (
	"fmt"
	"scratch/mockExample/util"
	"github.com/go-pg/pg"
)

// UserService ...
type UserService struct {
	DB *pg.DB
}

// NewUserService ...
func NewUserService() (*UserService, error) {

	db, err := util.DbConnect()
	if err != nil {
		fmt.Println("error connecting to db", err)
	}

	u := UserService{
		DB: db,
	}

	return &u, err
}

// CreateUpdateService ...
func (service *UserService) CreateUpdateService(req util.User, s Storage) (int, string, error) {

	var msg string

	result, err := s.CreateUpdate(req, *service.DB)
	if err != nil {
		fmt.Println("error with create/update", err)
		msg = err.Error()
		return 500, msg, err
	}

	msg = result

	return 200, msg, nil
}

// GetService ...
func (service *UserService) GetService(req *util.User, s *Store) (int, string, error) {

	var msg string

	result, err := s.Get(req, *service.DB)
	if err != nil {
		fmt.Println("error with get", err)
		if err.Error() == "pg: no rows in result set" {
			msg = "record not found"
			return 400, msg, err
		}

		msg = err.Error()
		return 500, msg, err
	}

	msg = result

	return 200, msg, nil
}

// DeleteService ...
func (service *UserService) DeleteService(req *util.User, s *Store) (int, string, error) {

	var msg string

	result, err := s.Delete(req, *service.DB)
	if err != nil {
		fmt.Println("error with delete", err)
		msg = err.Error()
		return 500, msg, err
	}

	msg = result

	return 200, msg, nil
}
