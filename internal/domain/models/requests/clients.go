package requests

import (
	"time"

	"github.com/TMg00000/customerscheduleapi/internal/domain/models/enums"
)

type Client struct {
	Id          int               `bson:"_id,omitempty"`
	Name        string            `bson:"name" json:"name"`
	PhoneNumber int               `bson:"phonenumber" json:"phonenumber"`
	TypeService enums.TypeService `bson:"typeservice" json:"typeservice"`
	DateTime    time.Time         `bson:"date,omitempty" json:"date,omitempty"`
}
