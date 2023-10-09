package skirpsionlineBE


// type T_Admin struct {
// 	AdmId			string 		`db:"adm_id" json:"adm_id"`
// 	AdmName			string		`db:"adm_name" json:"adm_name"`
// 	AdmUserName 	string		`db:"adm_username" json:"adm_username"`
// 	AdmPassWord 	string		`db:"adm_password" json:"adm_password"`
// 	AdmPhone 		string		`db:"adm_phone" json:"adm_phone"`
// 	AdmEmail 		string		`db:"adm_email" json:"adm_email"`
// 	AdmAddress		string		`db:"adm_address" json:"adm_address"`
// }

type B_ChekUser struct {
	UserName 	string	 	`json:"check_username"`
	PassWord 	string		`json:"check_password"`

}

// type InsertAdmin struct {
// 	InsertAdminBody		T_Admin	`json:"data"`
// }

type CheckUser struct {
	CheckUserBody		B_ChekUser	`json:"data"`
}

type CustomError struct {
    Code    int
    Message string
}


type B_Role struct {
	Role 		string	 	`json:"role"`
}