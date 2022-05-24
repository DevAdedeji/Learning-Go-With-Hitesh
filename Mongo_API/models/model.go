package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Neflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:""`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
