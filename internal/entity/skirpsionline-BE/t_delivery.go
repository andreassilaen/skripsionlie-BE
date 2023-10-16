package skirpsionlineBE

import "time"

type T_Delivery struct {
	ProdId    			string	`db:"prod_id" json:"prod_id"`
	AdmId				string	`db:"adm_id" json:"adm_id"`
	CtgId     			int 	`db:"ctg_id" json:"ctg_id"`
	DeliveryDate		time.Time	`db:"delivery_date" json:"delivery_date"`
}


type InsertDelivery struct {
	InsertDeliveryBody		T_Delivery	`json:"data"`
}
