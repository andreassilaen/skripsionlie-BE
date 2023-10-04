package skirpsionlineBE

type T_Product struct {
	ProdId    string `db:"prod_id" json:"prod_id"`
	CtgId     int `db:"ctg_id" json:"ctg_id"`
	ProdName  string `db:"prod_name" json:"prod_name"`
	ProdStock int `db:"prod_stock" json:"prod_stock"`
	ProdPrice int `db:"prod_price" json:"prod_price"`
	ProdDesc  string `db:"prod_desc" json:"prod_desc"`
}


type InsertProduct struct {
	InsertProductBody		T_Product	`json:"data"`
}
