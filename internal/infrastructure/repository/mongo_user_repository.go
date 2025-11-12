package repository

import (
	"context"
	"templateApiRestGo/internal/domain"
	"templateApiRestGo/internal/infrastructure/db"
)

// 🧠 Adaptador concreto que implementa la interfaz UserRepository.
// El servicio solo conoce la interfaz, no sabe que esto usa Mongo.

type MongoUserRepository struct {
	client *db.MongoClient
}

// Constructor
func NewMongoUserRepository(client *db.MongoClient) *MongoUserRepository {
	return &MongoUserRepository{client: client}
}

// Implementación concreta del método Create
func (r *MongoUserRepository) Create(user *domain.User) error {
	collection := r.client.Database.Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	return err
}
