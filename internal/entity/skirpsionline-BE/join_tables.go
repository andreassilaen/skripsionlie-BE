package skirpsionlineBE




type JoinAdmCust struct {
	CusData       	T_Customer `json:"data"`
	
	AdminData 		T_Admin `json:"data2"`
	

}


type InsertJoinHeaderDetailCart struct {
	HeaderCartBody		TH_Cart2	`json:"header"`

	DetailCartBody		[]TD_Cart2	`json:"detail"`


}

type JoinOrdCustTHTra struct {
	OrdId			int 		`db:"ord_id" json:"ord_id"`
	AdmId			string 		`db:"adm_id" json:"adm_id"`
	TranId    	string		`db:"tra_id" json:"tra_id"`
	
	CustName     string `db:"cust_name" json:"cust_name"`
	CustAddress  string `db:"cust_address" json:"cust_address"`

	TranTotal   string		`db:"tra_total" json:"tra_total"`
	OrdLastupdate	string	`db:"ord_lastupdate" json:"ord_lastupdate"`
}

type JoinTDTraProdByTraId struct {
	TraId             string `db:"tra_id" json:"tra_id"`


	// ProdId    			int		`db:"prod_id" json:"prod_id"`
	ProdName  			string 		`db:"prod_name" json:"prod_name"`

	TraDtlQty         int    `db:"tradtl_qty" json:"tradtl_qty"`
	TraDtlAmount		int		`db:"tradtl_amount" json:"tradtl_amount"`


}


