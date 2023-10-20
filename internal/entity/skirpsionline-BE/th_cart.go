package skirpsionlineBE

// import "time"

type TH_Cart struct {
	CartId     			string 		`db:"cart_id" json:"cart_id"`
	CustId  			string 		`db:"cust_id" json:"cust_id"`
	CartTotal			int 		`db:"cart_total" json:"cart_total"`
	CartLastupdate		string		`db:"cart_lastupdate" json:"cart_lastupdate"`
}


type TH_Cart2 struct {
	CustId  			string 		`db:"cust_id" json:"cust_id"`
	CartTotal			int 		`db:"cart_total" json:"cart_total"`
}


type InsertHeaderCart struct {
	InsertHeaderCartBody		TH_Cart2	`json:"data"`
}