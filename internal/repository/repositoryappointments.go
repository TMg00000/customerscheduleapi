package repository

import (
	"context"
	"time"

	"github.com/TMg00000/customerscheduleapi/internal/domain/models/requests"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppointmentsRepository struct {
	collection *mongo.Collection
}

func NewAppointmentsRepository(client *mongo.Client) *AppointmentsRepository {
	collection := client.Database("appointmentsdb").Collection("appointments")
	return &AppointmentsRepository{collection: collection}
}

func (r *AppointmentsRepository) Create(c requests.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, bson.M{
		"Name":        c.Name,
		"PhoneNumber": c.PhoneNumber,
		"TypeService": c.TypeService,
		"DateTime":    c.DateTime,
	})
	return err
}

func (r *AppointmentsRepository) GetAll() ([]requests.Client, error) {
	ctx := context.Background()
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var client []requests.Client
	if err = cursor.All(ctx, &client); err != nil {
		return nil, err
	}

	return client, nil
}

func (r *AppointmentsRepository) Update(c requests.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": c.Id}
	update := bson.M{
		"$set": bson.M{
			"Name":        c.Name,
			"PhoneNumber": c.PhoneNumber,
			"TypeService": c.TypeService,
			"DateTime":    c.DateTime,
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *AppointmentsRepository) Delete(Id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"Id": Id})
	return err
}
