package topup

import (
	"fmt"
	"time"
)

type TopUp struct {
	ID				int			`json:"id"`
	User_ID 		int			`json:"user_id"`
	Balance 		float64		`json:"balance"`
	Created_At 		time.Time	`json:"created_at"`
	Updated_At 		time.Time	`json:"updated_at"`

func topupPayment() {
	
}
