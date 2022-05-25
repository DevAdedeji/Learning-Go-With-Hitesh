package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Neflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie"`
	Watched bool               `json:"watched"`
}
