package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price"`
	Tags        []string           `bson:"tags" json:"tags"`
	Freelancer  primitive.ObjectID `bson:"freelancer" json:"freelancer"`
	City        string             `bson:"city" json:"city"`
	CreatedAt   int64              `bson:"created_at" json:"created_at"`
}
