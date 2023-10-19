package skirpsionlineBE

// type InsertSttkMJadwal struct {
// 	JadwalData SttkMJadwal ` json:"data"` //struct didalam struct
// }

// type InsertSttkEntriesHeader struct {
// 	EntriesHeaderData SttkEntriesHeader ` json:"data"` //struct didalam struct
// }

// type GetMasterJadwal struct {
// 	JwlPtid      int         `db:"jwl_ptid" json:"jwl_ptid"`
// 	JwlCode      string      `db:"jwl_code" json:"jwl_code"`
// 	JwlDateSche  string      `db:"jwl_date_sche" json:"jwl_date_sche"`
// 	JwlKwartal   float32     `db:"jwl_kwartal" json:"jwl_kwartal"`
// 	JwlComco     string      `db:"jwl_comco" json:"jwl_comco"`
// 	JwlTime      string      `db:"jwl_time" json:"jwl_time"`
// 	JwlNip       string      `db:"jwl_nip" json:"jwl_nip"`
// 	JwlNipName   zero.String `db:"jwl_nip_name" json:"jwl_nip_name"`
// 	JwlPic       string      `db:"jwl_pic" json:"jwl_pic"`
// 	JwlDateFinal zero.String `db:"jwl_date_final" json:"jwl_date_final"`
// 	JwlDateChg   string      `db:"jwl_date_chg" json:"jwl_date_chg"`
// 	JwlDateIsupd zero.String `db:"jwl_date_isupd" json:"jwl_date_isupd"`
// }

// // SttkMJadwal model
// type SttkMJadwal struct {
// 	JwlPtid             int         `db:"jwl_ptid" json:"jwl_ptid"`
// 	JwlCode             string      `db:"jwl_code" json:"jwl_code"`
// 	JwlComco            string      `db:"jwl_comco" json:"jwl_comco"`
// 	JwlDateSche         string      `db:"jwl_date_sche" json:"jwl_date_sche"`
// 	JwlDateSttk         time.Time   `db:"jwl_date_sttk" json:"jwl_date_sttk"`
// 	JwlKwartal          float32     `db:"jwl_kwartal" json:"jwl_kwartal"`
// 	JwlTime             string      `db:"jwl_time" json:"jwl_time"`
// 	JwlNip              string      `db:"jwl_nip" json:"jwl_nip"`
// 	JwlNipName          zero.String `db:"jwl_nip_name" json:"jwl_nip_name"`
// 	JwlPic              string      `db:"jwl_pic" json:"jwl_pic"`
// 	JwlPicName          string      `db:"jwl_pic_name" json:"jwl_pic_name"`
// 	JwlTeamCode         zero.String `db:"jwl_team_code" json:"jwl_team_code"`
// 	JwlIsupdate         zero.String `db:"jwl_isupdate" json:"jwl_isupdate"`
// 	JwlDateFinal        zero.String `db:"jwl_date_final" json:"jwl_date_final"`
// 	JwlDateIsupd        zero.String `db:"jwl_date_isupd" json:"jwl_date_isupd"`
// 	JwlNumber           string      `db:"jwl_number" json:"jwl_number"`
// 	JwlDateChg          string      `db:"jwl_date_chg" json:"jwl_date_chg"`
// 	JwlSttktype         string      `db:"jwl_sttktype" json:"jwl_sttktype"`
// 	JwlSelisihp         zero.String `db:"jwl_selisihp" json:"jwl_selisihp"`
// 	JwlSelisihm         zero.String `db:"jwl_selisihm" json:"jwl_selisihm"`
// 	JwlCrcawalp         zero.String `db:"jwl_crcawalp" json:"jwl_crcawalp"`
// 	JwlCrcawalm         zero.String `db:"jwl_crcawalm" json:"jwl_crcawalm"`
// 	JwlSales            zero.String `db:"jwl_sales" json:"jwl_sales"`
// 	JwlHakhil           zero.String `db:"jwl_hakhil" json:"jwl_hakhil"`
// 	JwlExtracash        zero.String `db:"jwl_extracash" json:"jwl_extracash"`
// 	JwlValue            zero.String `db:"jwl_value" json:"jwl_value"`
// 	JwlNewstore         int         `db:"jwl_newstore" json:"jwl_newstore"`
// 	JwlPayrflag         zero.String `db:"jwl_payflag" json:"jwl_payflag"`
// 	JwlPayrperiod       zero.String `db:"jwl_payperiod" json:"jwl_payperiod"`
// 	JwlPeriod           zero.String `db:"jwl_period" json:"jwl_period"`
// 	JwlHakhilflo        zero.String `db:"jwl_hakhilflo" json:"jwl_hakhilflo"`
// 	JwlJnssttk          string      `db:"jwl_jnssttk" json:"jwl_jnssttk"`
// 	JwlStatus           zero.String `db:"jwl_status" json:"jwl_status"`
// 	JwlValueEdpharos    float32     `db:"jwl_value_edpharos" json:"jwl_value_edpharos"`
// 	JwlValueEdnonpharos float32     `db:"jwl_value_ednonpharos" json:"jwl_value_ednonpharos"`
// }

type T_Customer struct {
	CustId       int `db:"cust_id" json:"cust_id"`
	CustName     string `db:"cust_name" json:"cust_name"`
	CustUserName string `db:"cust_username" json:"cust_username"`
	CustPassWord string `db:"cust_password" json:"cust_password"`
	CustPhone    string `db:"cust_phone" json:"cust_phone"`
	CustEmail    string `db:"cust_email" json:"cust_email"`
	CustAddress  string `db:"cust_address" json:"cust_address"`
}


type InsertCustomer struct {
	InsertCustomerBody		T_Customer	`json:"data"`
}

type T_Customer2 struct {
	// CustId       string `db:"cust_id" json:"cust_id"`
	CustName     string `db:"cust_name" json:"cust_name"`
	CustUserName string `db:"cust_username" json:"cust_username"`
	CustPassWord string `db:"cust_password" json:"cust_password"`
	CustPhone    string `db:"cust_phone" json:"cust_phone"`
	CustEmail    string `db:"cust_email" json:"cust_email"`
	CustAddress  string `db:"cust_address" json:"cust_address"`
}

type UpdateCustomerById struct {
	UpdateCustomerByIdBody		T_Customer2	`json:"data"`
}
