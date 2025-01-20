package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name"`
	Age    int                `bson:"age"`
	Course string             `bson:"course"`
	Roll   string             `bson:"roll"`
}
