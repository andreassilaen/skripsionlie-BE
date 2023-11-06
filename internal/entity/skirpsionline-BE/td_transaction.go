package skirpsionlineBE

// import "time"

type TD_Transaction struct {
	TraId             string `db:"tra_id" json:"tra_id"`
	ProdId             int    `db:"prod_id" json:"prod_id"`
	TraDtlQty         int    `db:"tradtl_qty" json:"tradtl_qty"`
	TraDtlPrice			int		`db:"tradtl_price" json:"tradtl_price"`
	TraDtlAmount		int		`db:"tradtl_amount" json:"tradtl_amount"`
}

type TD_Transaction2 struct {
	TraId             string `db:"tra_id" json:"tra_id"`
	ProdId             string    `db:"prod_id" json:"prod_id"`
	TraDtlQty         int    `db:"tradtl_qty" json:"tradtl_qty"`
	TraDtlPrice			float64		`db:"tradtl_price" json:"tradtl_price"`
	TraDtlAmount		float64		`db:"tradtl_amount" json:"tradtl_amount"`
}

type InsertDetailTransaction struct {
	InsertDetailTransactionBody TD_Transaction2 `json:"data"`
}
