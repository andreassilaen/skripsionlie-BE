package skirpsionlineBE

// import "time"

type TH_Transaction struct {
	TranId    	string		`db:"tra_id" json:"tra_id"`
	CartId    	string		`db:"cart_id" json:"cart_id"`
	CustId 		string		`db:"cust_id" json:"cust_id"`
	RekId     	string `db:"rek_id" json:"rek_id"`
	TranTotal   int		`db:"tra_total" json:"tra_total"`
	TranImg    	string		`db:"tra_img" json:"tra_img"`
	TraDate		string	`db:"tra_date" json:"tra_date"`
	TraCheckedYN	string	`db:"tra_checkedyn" json:"tra_checkedyn"`
}


// type InsertTransactionHeader struct {
// 	InsertTransactionDetailBody		TH_Transaction	`json:"data"`
// }


type TH_Transaction2 struct {
	CartId    	string		`db:"cart_id" json:"cart_id"`
	CustId 		string		`db:"cust_id" json:"cust_id"`
	RekId     	string 		`db:"rek_id" json:"rek_id"`
	TranTotal   int		`db:"tra_total" json:"tra_total"`
	TranImg    	string		`db:"tra_img" json:"tra_img"`
}

type InsertHeaderTransaction struct {
	InsertHeaderTransactionBody		TH_Transaction2	`json:"data"`
}