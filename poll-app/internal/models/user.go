package models

import (
	"context"
	"poll-app/ent"
	"poll-app/ent/user"
	"time"
)

// User represents a user entity in the system
type User struct {
	ID        int       `json:"id,omitempty"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateUser creates a new user with the given username and password
func CreateUser(ctx context.Context, username, password string) error {
	client := ctx.Value("client").(*ent.Client)
	_, err := client.User.
		Create().
		SetUsername(username).
		SetPassword(password).
		SetCreatedAt(time.Now()).
		Save(ctx)
	return err
}

// GetUserByUsername retrieves a user by their username
func GetUserByUsername(ctx context.Context, username string) (*User, error) {
	client := ctx.Value("client").(*ent.Client)
	u, err := client.User.
		Query().
		Where(user.UsernameEQ(username)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
	}, nil
}
