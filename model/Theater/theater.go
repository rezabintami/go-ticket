package theater

import "time"

type Theater struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Place      string    `json:"place"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
	Deleted_At time.Time `json:"deleted_at"`
}
