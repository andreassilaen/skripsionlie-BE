package skirpsionlineBE




type JoinAdmCust struct {
	CusData       	T_Customer `json:"data"`
	
	AdminData 		T_Admin `json:"data2"`
	

}


type InsertJoinHeaderDetailCart struct {
	HeaderCartBody		TH_Cart2	`json:"header"`

	DetailCartBody		[]TD_Cart2	`json:"detail"`


}
