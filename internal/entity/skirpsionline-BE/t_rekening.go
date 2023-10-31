package skirpsionlineBE


type T_Rekening struct {
	RekId     	int `db:"rek_id" json:"rek_id"`
	RekBank  	string `db:"rek_bank" json:"rek_bank"`
	RekNumber  	int `db:"rek_number" json:"rek_number"`
	RekName  	string `db:"rek_name" json:"rek_name"`
}

