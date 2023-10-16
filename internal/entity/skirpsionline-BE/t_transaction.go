package skirpsionlineBE

// import "time"

type T_Transaction struct {
	TranId    	string		`db:"tra_id" json:"tra_id"`
	CartId    	string		`db:"cart_id" json:"cart_id"`
	TranTotal   string		`db:"tra_total" json:"tra_total"`
	TranImg    	string		`db:"tra_img" json:"tra_img"`
	TraDate		string	`db:"tra_date" json:"tra_date"`
}


type InsertTransaction struct {
	InsertTransactionBody		T_Transaction	`json:"data"`
}
