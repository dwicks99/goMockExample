package userservice

import (
	"fmt"
	"scratch/mockExample/util"
	"testing"
)

func TestUserService_CreateUpdate(t *testing.T) {
	tests := []struct {
		name             string
		user             util.User
		wantCode         int
		wantStringResult string
		wantErr          bool
		store            *util.MockStore
	}{
		{
			name:             "success",
			user:             util.User{Name: "Banny Wicks", Email: "bwicks@us.ci.org"},
			wantCode:         200,
			wantStringResult: "...successfully inserted or updated Banny Wicks bwicks@us.ci.org",
			wantErr:          false,
			store:            &util.MockStore{},
		},
				{
			name:             "error",
			user:             util.User{Name: "Banny Wicks", Email: "bwicks@us.ci.org"},
			wantCode:         500,
			wantStringResult: "internal server error",
			wantErr:          true,
			store:            &util.MockStore{
				CreateUpdateError: "error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			myservice, err := NewUserService()
			if err != nil {
				fmt.Println(err)
				return
			}

			statusCode, stringResult, err := myservice.CreateUpdateService(tt.user, tt.store)
			
			if !tt.wantErr {
				if err != nil {
				t.Errorf("UserService.CreateUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
				}
			} else {
				if err == nil{
					t.Errorf("Expected to get error and did not")
					return
				}
			}
			
			if statusCode != tt.wantCode {
				t.Errorf("UserService.CreateUpdate() status code got %v want %v", statusCode, tt.wantCode)
			}
			if stringResult != tt.wantStringResult {
				t.Errorf("UserService.CreateUpdate() string result got %v want %v", stringResult, tt.wantStringResult)
			}
			if tt.store.CreateUpdateCalledTimes != 1 {
				t.Errorf("expected CreateUpdateCalledTime to be 1, got %v", tt.store.CreateUpdateCalledTimes)
			}
			if tt.user.Name != tt.store.CreateUpdateCalledWithName {
				t.Errorf("expected CreateUpdateCalledWithName to be %v, got %v", tt.user.Name, tt.store.CreateUpdateCalledWithName)
			}
			if tt.user.Email != tt.store.CreateUpdateCalledWithEmail {
				t.Errorf("expected CreateUpdateCalledWithEmail to be %v, got %v", tt.user.Email, tt.store.CreateUpdateCalledWithEmail)
			}
		})
	}
}
