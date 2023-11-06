package skirpsionlineBE

// import "time"

type T_Product struct {
	ProdId    			int		`db:"prod_id" json:"prod_id"`
	AdmId				string		`db:"adm_id" json:"adm_id"`
	CtgId     			int 		`db:"ctg_id" json:"ctg_id"`
	CtgType    			string 		`db:"ctg_type" json:"ctg_type"`
	ProdName  			string 		`db:"prod_name" json:"prod_name"`
	ProdDesc  			string 		`db:"prod_desc" json:"prod_desc"`
	ProdPrice 			int 		`db:"prod_price" json:"prod_price"`
	ProdStock 			int 		`db:"prod_stock" json:"prod_stock"`
	ProdLastupdate		string		`db:"prod_lastupdate" json:"prod_lastupdate"`
	ProdImage			string		`db:"prod_img" json:"prod_img"`
}

type T_Product2 struct {
	// ProdId    			int		`db:"prod_id" json:"prod_id"`
	AdmId				string		`db:"adm_id" json:"adm_id"`
	CtgId     			int 		`db:"ctg_id" json:"ctg_id"`
	ProdName  			string 		`db:"prod_name" json:"prod_name"`
	ProdDesc  			string 		`db:"prod_desc" json:"prod_desc"`
	ProdPrice 			int 		`db:"prod_price" json:"prod_price"`
	ProdStock 			int 		`db:"prod_stock" json:"prod_stock"`
	ProdImage			string		`db:"prod_img" json:"prod_img"`
	// ProdLastupdate		string	`db:"prod_lastupdate" json:"prod_lastupdate"`
}

type T_Product3 struct {
	// ProdId    			int		`db:"prod_id" json:"prod_id"`
	AdmId				string		`db:"adm_id" json:"adm_id"`
	ProdName  			string 		`db:"prod_name" json:"prod_name"`
	ProdDesc  			string 		`db:"prod_desc" json:"prod_desc"`
	ProdPrice 			int 		`db:"prod_price" json:"prod_price"`
	ProdStock 			int 		`db:"prod_stock" json:"prod_stock"`
	ProdImage			string		`db:"prod_img" json:"prod_img"`
	// ProdLastupdate		string	`db:"prod_lastupdate" json:"prod_lastupdate"`
}


type InsertProduct struct {
	InsertProductBody		T_Product2	`json:"data"`
}

type UpdateProdById struct {
	UpdateProdByIdBody		T_Product3	`json:"data"`
}