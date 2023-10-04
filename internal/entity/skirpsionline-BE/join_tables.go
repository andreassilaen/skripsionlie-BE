package skirpsionlineBE

type JoinAdmCust struct {
	CustId       string `db:"cust_id" json:"cust_id"`
	CustUserName string `db:"cust_username" json:"cust_username"`
	CustPassWord string `db:"cust_password" json:"cust_password"`
	
	AdmId			string 		`db:"adm_id" json:"adm_id"`
	AdmUserName 	string		`db:"adm_username" json:"adm_username"`
	AdmPassWord 	string		`db:"adm_password" json:"adm_password"`

}
