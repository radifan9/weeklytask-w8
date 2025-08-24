package utils

import (
	"errors"
	"fmt"
)

type User struct {
	ID   string
	Name string
}

// NewUser creates a new User instance
func NewUser(id, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}

// UserManager manages a collection of users
// key : ID, value : pointer to user instance
type UserManager struct {
	users map[string]*User
}

// NewUserManager creates and initializes a new UserManager
// The users map must be initialized to avoid nil map panics
func NewUserManager() UserManager {
	return UserManager{
		// Initialize the users map, when UserManager is created
		users: make(map[string]*User),
	}
}

// AddUser adds a new user to the manager
func (um *UserManager) AddUser(newId, newName string) (string, error) {
	if user, exists := um.users[newId]; exists {
		errorMsg := fmt.Sprintf("failed, id:%s already registered", user.ID)
		return "", errors.New(errorMsg)
	}
	um.users[newId] = &User{ID: newId, Name: newName}
	successMsg := fmt.Sprintf("success, added id:%s to DB", newId)
	return successMsg, nil
}

// GetUser retrieves a user by ID
func (um UserManager) GetUser(id string) (string, error) {
	if user, exists := um.users[id]; exists {
		userData := fmt.Sprintf("id: %s\nName: %s", user.ID, user.Name)
		return userData, nil
	} else {
		// If ID doesn't exist
		errorMsg := fmt.Sprintf("failed, cannot find id: %s", id)
		return "", errors.New(errorMsg)
	}
}
