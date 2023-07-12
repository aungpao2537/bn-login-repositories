package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type FirebaseRepository struct {
	Context    context.Context
	Collection *mongo.Collection
}

type IFirebaseRepository interface {
}
