package repository

import (
	"context"
	"time"

	"github.com/TMg00000/customerscheduleapi/internal/domain/models/requests"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppointmentsRepository struct {
	colletion *mongo.Collection
}

func NewAppointmentsRepository(client *mongo.Client) *AppointmentsRepository {
	colletion := client.Database("appointmentsdb").Collection("appointments")
	return &AppointmentsRepository{colletion: colletion}
}

func (r *AppointmentsRepository) Create(c requests.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.colletion.InsertOne(ctx, bson.M{
		"Name":        c.Name,
		"PhoneNumber": c.PhoneNumber,
		"TypeService": c.TypeService,
		"DateTime":    c.DateTime,
	})
	return err
}

func (r *AppointmentsRepository) GetAll() ([]requests.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.colletion.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var client []requests.Client
	for cursor.Next(ctx) {
		var c requests.Client
		if err := cursor.Decode(&c); err != nil {
			return nil, err
		}
		client = append(client, c)
	}

	return client, nil
}

func (r *AppointmentsRepository) Update(c requests.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": c.Id}
	update := bson.M{
		"Name":        c.Name,
		"PhoneNumber": c.PhoneNumber,
		"TypeService": c.TypeService,
		"DateTime":    c.DateTime,
	}

	_, err := r.colletion.UpdateOne(ctx, filter, update)
	return err
}

func (r *AppointmentsRepository) Delete(Id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.colletion.DeleteOne(ctx, bson.M{"Id": Id})
	return err
}
