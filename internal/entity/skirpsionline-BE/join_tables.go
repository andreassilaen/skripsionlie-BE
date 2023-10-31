package skirpsionlineBE

type JoinAdmCust struct {
	CusData T_Customer `json:"data"`

	AdminData T_Admin `json:"data2"`
}

type InsertJoinHeaderDetailCart struct {
	HeaderCartBody TH_Cart2 `json:"header"`

	DetailCartBody []TD_Cart2 `json:"detail"`
}

type JoinOrdCustTHTra struct {
	OrdId  int    `db:"ord_id" json:"ord_id"`
	AdmId  string `db:"adm_id" json:"adm_id"`
	TranId string `db:"tra_id" json:"tra_id"`

	CustName    string `db:"cust_name" json:"cust_name"`
	CustAddress string `db:"cust_address" json:"cust_address"`

	TranTotal     string `db:"tra_total" json:"tra_total"`
	OrdLastupdate string `db:"ord_lastupdate" json:"ord_lastupdate"`
}

type JoinTDTraProdByTraId struct {
	TraId string `db:"tra_id" json:"tra_id"`

	// ProdId    			int		`db:"prod_id" json:"prod_id"`
	ProdName string `db:"prod_name" json:"prod_name"`

	TraDtlQty    int `db:"tradtl_qty" json:"tradtl_qty"`
	TraDtlAmount int `db:"tradtl_amount" json:"tradtl_amount"`
}

// type JoinTHTDCart struct {
// 	CartId     			string 		`db:"cart_id" json:"cart_id"`
// 	CustId  			string 		`db:"cust_id" json:"cust_id"`
// 	CartTotal			int 		`db:"cart_total" json:"cart_total"`
// 	CartPayedYN			string		`db:"cart_payedyn" json:"cart_payedyn"`
// }

type JoinTHTDCartProd struct {
	CartId      string `db:"cart_id" json:"cart_id"`
	CustId      string `db:"cust_id" json:"cust_id"`
	CartTotal   int    `db:"cart_total" json:"cart_total"`
	CartPayedYN string `db:"cart_payedyn" json:"cart_payedyn"`

	ProdId    string `db:"prod_id" json:"prod_id"`
	ProdName  string `db:"prod_name" json:"prod_name"`
	ProdDesc  string `db:"prod_desc" json:"prod_desc"`
	ProdPrice int    `db:"prod_price" json:"prod_price"`
	ProdStock int    `db:"prod_stock" json:"prod_stock"`

	CartDtlQty int `db:"cardtl_qty" json:"cardtl_qty"`

	ProdImage string `db:"prod_img" json:"prod_img"`
}

type JoinTHTDCartProd2 struct {
	CartDtlQty int    `db:"cardtl_qty" json:"cardtl_qty"`
	CustId     string `db:"cust_id" json:"cust_id"`
	CartId     string `db:"cart_id" json:"cart_id"`
	ProdId     string `db:"prod_id" json:"prod_id"`
}

type UpdateQtyDetailJoinTHTDCartProd struct {
	UpdateQtyDetailJoinTHTDCartProdBody JoinTHTDCartProd2 `json:"data"`
}
