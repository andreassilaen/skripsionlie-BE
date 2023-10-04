package skirpsionlineBE

import "time"

type TH_Order struct {
	OrdId     	string `db:"ord_id" json:"ord_id"`
	CustomerId  string `db:"customer_id" json:"customer_id"`
	OrdStatus	string `db:"ord_status" json:"ord_status"`
	OrdTotal	string `db:"ord_total" json:"ord_total"`
	OrdPayment	string `db:"ord_payment" json:"ord_payment"`
	OrdDate		time.Time `db:"ord_date" json:"ord_date"`
}
