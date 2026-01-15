package requests

import (
	"time"

	"github.com/TMg00000/customerscheduleapi/internal/domain/enums"
)

type Client struct {
	name        string            `json:"name"`
	phoneNumber int               `json:"phonenumber"`
	typeService enums.TypeService `json:"typeservice"`
	dateTime    time.Time         `json:"date,omitempty"`
}
