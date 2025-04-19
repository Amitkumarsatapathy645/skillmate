package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Skills      []string           `bson:"skills" json:"skills"`
	Budget      float64            `bson:"budget" json:"budget"`
	City        string             `bson:"city" json:"city"`
	Client      primitive.ObjectID `bson:"client" json:"client"`
	CreatedAt   int64              `bson:"created_at" json:"created_at"`
}
