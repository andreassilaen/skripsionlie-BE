package skirpsionlineBE

type T_Product struct {
	ProdId    string `db:"prod_id" json:"prod_id"`
	CtgId     string `db:"ctg_id" json:"ctg_id"`
	ProdName  string `db:"prod_name" json:"prod_name"`
	ProdStock string `db:"prod_stock" json:"prod_stock"`
	ProdPrice string `db:"prod_price" json:"prod_price"`
	ProdDesc  string `db:"prod_desc" json:"prod_desc"`
}
