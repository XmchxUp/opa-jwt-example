package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	QueryByID(ctx context.Context, userID uuid.UUID) (User, error)
}

// Core manages the set of APIs for user access.
type Core struct {
	storer Storer
}

func NewCore(storer Storer) *Core {
	return &Core{storer: storer}
}

// QueryByID finds the user by the specified ID.
func (c *Core) QueryByID(ctx context.Context, userID uuid.UUID) (User, error) {
	user, err := c.storer.QueryByID(ctx, userID)
	if err != nil {
		return User{}, fmt.Errorf("query: userID[%s]: %w", userID, err)
	}

	return user, nil
}
