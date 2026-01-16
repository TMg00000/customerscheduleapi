package repository

import (
	"context"
	"time"

	"github.com/TMg00000/customerscheduleapi/internal/domain/models/requests"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		"Name":        c.Name,
		"PhoneNumber": c.PhoneNumber,
		"TypeService": c.TypeService,
		"DateTime":    c.DateTime,
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *AppointmentsRepository) Delete(Id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"Id": Id})
	return err
}

func GetNextId(collection *mongo.Collection, ctx context.Context) (int, error) {
	filter := bson.M{"_id": "appointmentId"}
	update := bson.M{"$inc": bson.M{"seq": 1}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)

	var result struct {
		Seq int `bson:"seq"`
	}

	err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return 0, err
	}
	return result.Seq, nil
}
