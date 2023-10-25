package skirpsionlineBE

// import "time"

type TD_Cart struct {
	CartId             string `db:"cart_id" json:"cart_id"`
	ProdId             int    `db:"prod_id" json:"prod_id"`
	CartDtlQty         int    `db:"cardtl_qty" json:"cardtl_qty"`
	CartDtlConfirmedYN int    `db:"cardtl_confirmedyn" json:"cardtl_confirmedyn"`
}

type TD_Cart2 struct {
	CartId     string `db:"cart_id" json:"cart_id"`
	ProdId     string `db:"prod_id" json:"prod_id"`
	CartDtlQty int    `db:"cardtl_qty" json:"cardtl_qty"`
}

type InsertDetailCart struct {
	InsertDetailCartBody TD_Cart2 `json:"data"`
}
