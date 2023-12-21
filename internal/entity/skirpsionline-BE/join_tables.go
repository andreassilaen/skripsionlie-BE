package skirpsionlineBE

type JoinAdmCust struct {
	CusData T_Customer `json:"data"`

	AdminData T_Admin `json:"data2"`
}

type InsertJoinHeaderDetailCart struct {
	HeaderCartBody TH_Cart2 `json:"header"`

	DetailCartBody TD_Cart2 `json:"detail"`
	// DetailCartBody []TD_Cart2 `json:"detail"`
}

type InsertJoinHeaderDetailTran struct {
	HeaderTranBody TH_Transaction2 `json:"header"`

	DetailTranBody []TD_Transaction2 `json:"detail"`
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

	CustName    string `db:"cust_name" json:"cust_name"`
	CustAddress string `db:"cust_address" json:"cust_address"`
}

type JoinOrdTHTDTraProdByOrdId struct {
	OrdId string `db:"ord_id" json:"ord_id"`

	TraId string `db:"tra_id" json:"tra_id"`

	// ProdId    			int		`db:"prod_id" json:"prod_id"`
	ProdName string `db:"prod_name" json:"prod_name"`

	TraDtlQty    int `db:"tradtl_qty" json:"tradtl_qty"`
	TraDtlAmount int `db:"tradtl_amount" json:"tradtl_amount"`

	CustName    string `db:"cust_name" json:"cust_name"`
	CustAddress string `db:"cust_address" json:"cust_address"`
}

type JoinTDTranProdCustByTraId struct {
	TraId string `db:"tra_id" json:"tra_id"`

	ProdId       int    `db:"prod_id" json:"prod_id"`
	ProdName     string `db:"prod_name" json:"prod_name"`
	ProdPrice    int    `db:"prod_price" json:"prod_price"`
	ProdStock    int    `db:"prod_stock" json:"prod_stock"`
	TraDtlQty    int    `db:"tradtl_qty" json:"tradtl_qty"`
	TraDtlAmount int    `db:"tradtl_amount" json:"tradtl_amount"`
	TraTotal     int    `db:"tra_total" json:"tra_total"`

	CustName    string `db:"cust_name" json:"cust_name"`
	CustAddress string `db:"cust_address" json:"cust_address"`
}

type JoinTHTraRek struct {
	TraId int `db:"tra_id" json:"tra_id"`
	// CartId    	string		`db:"cart_id" json:"cart_id"`
	CustId       string `db:"cust_id" json:"cust_id"`
	RekBank      string `db:"rek_bank" json:"rek_bank"`
	TraTotal     int    `db:"tra_total" json:"tra_total"`
	TraImg       string `db:"tra_img" json:"tra_img"`
	TraDate      string `db:"tra_date" json:"tra_date"`
	TraCheckedYN string `db:"tra_checkedyn" json:"tra_checkedyn"`
}

type JoinOrdTHTraDelByCustId struct {
	OrdId string `db:"ord_id" json:"ord_id"`

	TraId    string `db:"tra_id" json:"tra_id"`
	TraTotal int    `db:"tra_total" json:"tra_total"`

	OrdConfirmedYN  string `db:"ord_confirmedyn" json:"ord_confirmedyn"`
	OrdOnDeliveryYN string `db:"ord_ondeliveryyn" json:"ord_ondeliveryyn"`
	OrdLastupdate   string `db:"ord_lastupdate" json:"ord_lastupdate"`

	DeliveryDoneYN string `db:"delivery_doneyn" json:"delivery_doneyn"`
	DeliveryDate string `db:"delivery_date" json:"delivery_date"`
}
type JoinOrdTHTraByCustId struct {
	OrdId string `db:"ord_id" json:"ord_id"`

	TraId    string `db:"tra_id" json:"tra_id"`
	TraTotal int    `db:"tra_total" json:"tra_total"`

	OrdConfirmedYN  string `db:"ord_confirmedyn" json:"ord_confirmedyn"`
	OrdOnDeliveryYN string `db:"ord_ondeliveryyn" json:"ord_ondeliveryyn"`
	OrdLastupdate   string `db:"ord_lastupdate" json:"ord_lastupdate"`
}

type CountTHTraOrdDel struct {
	TraUncheck  string `db:"tra_uncheck" json:"tra_uncheck"`
	OrdCanceled string `db:"ord_canceled" json:"ord_canceled"`
	OrdProcess  string `db:"ord_process" json:"ord_process"`
	OrdDelivery string `db:"ord_delivery" json:"ord_delivery"`
	DelOnGoing  string `db:"del_ongoing" json:"del_ongoing"`
	DelDoned    string `db:"del_doned" json:"del_doned"`
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

type JoinReportOrdTHTra struct {
	OrdId string `db:"ord_id" json:"ord_id"`

	TraId    string `db:"tra_id" json:"tra_id"`
	TraTotal int    `db:"tra_total" json:"tra_total"`

	OrdLastupdate string `db:"ord_lastupdate" json:"ord_lastupdate"`

	// OrdConfirmedYN 	string		`db:"ord_confirmedyn" json:"ord_confirmedyn"`
	// OrdOnDeliveryYN 	string		`db:"ord_ondeliveryyn" json:"ord_ondeliveryyn"`

	// DeliveryDoneYN 	string `db:"delivery_doneyn" json:"delivery_doneyn"`

}

type JoinReportOrdTHTraDel struct {
	OrdId string `db:"ord_id" json:"ord_id"`

	TraId    string `db:"tra_id" json:"tra_id"`
	TraTotal int    `db:"tra_total" json:"tra_total"`

	OrdLastupdate string `db:"ord_lastupdate" json:"ord_lastupdate"`

	DeliveryDoneYN 	string `db:"delivery_doneyn" json:"delivery_doneyn"`
	DeliveryDate 	string `db:"delivery_date" json:"delivery_date"`
	DeliveryImg 	string `db:"delivery_img" json:"delivery_img"`



}

type JoinDetailReport struct {
	OrdId         string `db:"ord_id" json:"ord_id"`
	OrdLastupdate string `db:"ord_lastupdate" json:"ord_lastupdate"`

	AdmName string `db:"adm_name" json:"adm_name"`

	TraId    string `db:"tra_id" json:"tra_id"`
	TraTotal int    `db:"tra_total" json:"tra_total"`

	CustName    string `db:"cust_name" json:"cust_name"`
	CustPhone   string `db:"cust_phone" json:"cust_phone"`
	CustAddress string `db:"cust_address" json:"cust_address"`

	EmpName      string `db:"emp_name" json:"emp_name"`
	DeliveryDate string `db:"delivery_date" json:"delivery_date"`

	// OrdConfirmedYN 	string		`db:"ord_confirmedyn" json:"ord_confirmedyn"`
	// OrdOnDeliveryYN 	string		`db:"ord_ondeliveryyn" json:"ord_ondeliveryyn"`

	// DeliveryDoneYN 	string `db:"delivery_doneyn" json:"delivery_doneyn"`

}
