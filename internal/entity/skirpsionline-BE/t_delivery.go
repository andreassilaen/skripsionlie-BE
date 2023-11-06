package skirpsionlineBE

type T_Delivery struct {
	EmpId string `db:"emp_id" json:"emp_id"`
	OrdId string `db:"ord_id" json:"ord_id"`
	// CtgId        int       `db:"ctg_id" json:"ctg_id "`
	DeliveryDoneYN 	string `db:"delivery_doneyn" json:"delivery_doneyn"`
	DeliveryDate string `db:"delivery_date" json:"delivery_date"`
}

type T_Delivery2 struct {
	EmpId string `db:"emp_id" json:"emp_id"`
	OrdId string `db:"ord_id" json:"ord_id"`
}

type InsertDelivery struct {
	InsertDeliveryBody T_Delivery2 `json:"data"`
}
