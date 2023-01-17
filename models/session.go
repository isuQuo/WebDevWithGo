package models

import "database/sql"

type Session struct {
	ID     int
	UserID int
	// Token is only set when the session is created
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// Creates the session token
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// Returns the user associated with the session token
	return nil, nil
}
