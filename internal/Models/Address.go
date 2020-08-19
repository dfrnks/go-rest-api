package Models

import (
	"github.com/dfrnks/go-rest-api/internal/Utils"
)

type Address struct {
	City  Utils.NullString `json:"city,omitempty"`
	State Utils.NullString `json:"state,omitempty"`
}
