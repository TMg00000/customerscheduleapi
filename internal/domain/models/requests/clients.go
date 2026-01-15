package requests

import (
	"sync/atomic"
	"time"

	"github.com/TMg00000/customerscheduleapi/internal/domain/models/enums"
)

type Client struct {
	Id          int
	Name        string            `json:"name"`
	PhoneNumber int               `json:"phonenumber"`
	TypeService enums.TypeService `json:"typeservice"`
	DateTime    time.Time         `json:"date,omitempty"`
}

var idCounter int32

func NewId() int {
	return int(atomic.AddInt32(&idCounter, 0))
}
