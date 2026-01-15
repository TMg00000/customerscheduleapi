package requests

import (
	"time"

	"github.com/TMg00000/customerscheduleapi/internal/domain/models/enums"
)

type UpdateClient struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	PhoneNumber int               `json:"phonenumber"`
	TypeService enums.TypeService `json:"typeservice"`
	DateTime    time.Time         `json:"date,omitempty"`
}
