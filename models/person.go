package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Person struct {
	Name string `json:"name" bson:"name"`
	Age  int    `json:"age" bson:"age"`
}

type PersonRepository interface {
	Create(ctx context.Context, person *Person) error
	GetAll(ctx context.Context) ([]Person, error)
}

type personRepository struct {
	collection *mongo.Collection
}

func NewPersonRepository(client *mongo.Client) PersonRepository {
	collection := client.Database("marrywith").Collection("persons")
	return &personRepository{collection}
}

func (r *personRepository) Create(ctx context.Context, person *Person) error {
	_, err := r.collection.InsertOne(ctx, person)
	return err
}

func (r *personRepository) GetAll(ctx context.Context) ([]Person, error) {
	cur, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var persons []Person
	for cur.Next(ctx) {
		var person Person
		err := cur.Decode(&person)
		if err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return persons, nil
}
