package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type firebaseRepository struct {
	Context    context.Context
	Collection *mongo.Collection
}

type IfirebaseRepository interface {
}
