package skirpsionlineBE

// import "time"

type T_Product struct {
	ProdId    			string		`db:"prod_id" json:"prod_id"`
	AdmId				string		`db:"adm_id" json:"adm_id"`
	CtgId     			int 		`db:"ctg_id" json:"ctg_id"`
	ProdName  			string 		`db:"prod_name" json:"prod_name"`
	ProdDesc  			string 		`db:"prod_desc" json:"prod_desc"`
	ProdPrice 			int 		`db:"prod_price" json:"prod_price"`
	ProdStock 			int 		`db:"prod_stock" json:"prod_stock"`
	ProdLastupdate		string	`db:"prod_lastupdate" json:"prod_lastupdate"`
}

type T_Product2 struct {
	ProdId    			string		`db:"prod_id" json:"prod_id"`
	AdmId				string		`db:"adm_id" json:"adm_id"`
	CtgId     			int 		`db:"ctg_id" json:"ctg_id"`
	ProdName  			string 		`db:"prod_name" json:"prod_name"`
	ProdDesc  			string 		`db:"prod_desc" json:"prod_desc"`
	ProdPrice 			int 		`db:"prod_price" json:"prod_price"`
	ProdStock 			int 		`db:"prod_stock" json:"prod_stock"`
	// ProdLastupdate		string	`db:"prod_lastupdate" json:"prod_lastupdate"`
}


type InsertProduct struct {
	InsertProductBody		T_Product2	`json:"data"`
}
