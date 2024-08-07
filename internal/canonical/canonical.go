package canonical

import (
	"time"
)

type Order struct {
	Id        string    `bson:"_id"`
	Products  []string  `bson:"products"`
	Total     float32   `bson:"total"`
	Status    string    `bson:"status"`
	CreatedAt time.Time `bson:"created_at"`
}
