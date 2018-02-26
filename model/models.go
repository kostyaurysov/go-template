package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// User type represents the registered user.

	User struct {
		ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"firstname"`
		LastName     string        `json:"lastname"`
		Email        string        `json:"email"`
		HashPassword []byte        `json:"hashpassword,omitempty"`
	}
)
