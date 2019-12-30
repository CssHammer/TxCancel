package wrapper

import (
	"txCancel/models/db"
	"txCancel/models/transport"
)

func WrapUser(user *db.User) *transport.User {
	return &transport.User{
		ID:      user.ID,
		Name:    user.Name,
		Balance: user.Balance,
	}
}
