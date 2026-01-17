package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/TMg00000/customerscheduleapi/internal/domain/models/requests"
	"github.com/TMg00000/customerscheduleapi/internal/domain/models/resources/resourceserrorsmessages"
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
		"name":        c.Name,
		"phonenumber": c.PhoneNumber,
		"typeservice": c.TypeService,
		"datetime":    c.DateTime,
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

	filter := bson.M{"_id": c.Id}
	update := bson.M{
		"$set": bson.M{
			"name":        c.Name,
			"phonenumber": c.PhoneNumber,
			"typetervice": c.TypeService,
			"datetime":    c.DateTime,
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *AppointmentsRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf(resourceserrorsmessages.ErrorDeleteAppointment)
	}

	return err
}
