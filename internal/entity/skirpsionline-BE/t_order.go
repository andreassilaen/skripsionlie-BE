package skirpsionlineBE

// import "time"

type T_Order struct {
	OrdId			string 		`db:"ord_id" json:"ord_id"`
	AdmId			string 		`db:"adm_id" json:"adm_id"`
	TraId			string 		`db:"tra_id" json:"tra_id"`
	OrdCancelYN 	string		`db:"ord_cancelyn" json:"ord_cancelyn"`
	OrdLastupdate	string	`db:"ord_lastupdate" json:"ord_lastupdate"`
}


type InsertOrder struct {
	InsertOrderBody		T_Order	`json:"data"`
}