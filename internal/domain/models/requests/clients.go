package requests

import (
	"time"

	"github.com/TMg00000/customerscheduleapi/internal/domain/models/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	PhoneNumber int                `bson:"phonenumber" json:"phonenumber"`
	TypeService enums.TypeService  `bson:"typeservice" json:"typeservice"`
	DateTime    time.Time          `bson:"date,omitempty" json:"date,omitempty"`
}
