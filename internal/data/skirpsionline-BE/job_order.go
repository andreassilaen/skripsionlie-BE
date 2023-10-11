package skirpsionlineBE

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"

	jaegerLog "skripsi-online-BE/pkg/log"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		stmt *map[string]*sqlx.Stmt

		tracer opentracing.Tracer
		logger jaegerLog.Factory
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

// Tambahkan query di dalam const
// getAllUser = "GetAllUser"
// qGetAllUser = "SELECT * FROM users"

const (
	////__________________________________________ T_Admin ____________________________________________
	//	+ selesai +
	getAllAdmin  = "GetAllAdmin"
	qGetAllAdmin = `
	SELECT adm_id,
		adm_name,
		adm_username,
		adm_password,
		adm_phone,
		adm_email,
		adm_address 
	FROM t_admin`

	getAdmLastData = "GetAdmLastData"
	qGetAdmLastData = `
	SELECT * FROM t_admin
	ORDER BY adm_id DESC
	LIMIT 1`

	insertAdmin = "InsertAdmin"
	qInsertAdmin = `
	INSERT INTO t_admin
		(adm_id,
		adm_name,
		adm_username,
		adm_password,
		adm_phone,
		adm_email,
		adm_address)
	VALUES(?, ?, ?, ?, ?,  ?, ?)`



	getAdmByLogin = "GetAdmByLogin"
	qGetAdmByLogin = `
	SELECT adm_id,
		adm_name,
		adm_username,
		adm_password,
		adm_phone,
		adm_email,
		adm_address 
	FROM t_admin
	WHERE  adm_username = ?
	AND adm_password = ?`

	////__________________________________________ T_Customer ____________________________________________
	//  - belum -
	getAllCustomer  = "GetAllCustomer"
	qGetAllCustomer = `
	SELECT cust_id,
		cust_name,
		cust_username,
		cust_password,
		cust_phone,
		cust_email,
		cust_address 
	FROM t_customer`

	getCustById= "GetCustById"
	qGetCustById = `
	SELECT cust_id,
		cust_name,
		cust_username,
		cust_password,
		cust_phone,
		cust_email,
		cust_address 
	FROM t_customer
	WHERE  
	cust_id = ?`

	getCountCust = "GetCountCust"
	qGetCountCust = `
	SELECT COUNT(cust_id) AS total
	FROM t_customer`

	getCustLastData = "GetCustLastData"
	qGetCustLastData = `
	SELECT * FROM t_customer
	ORDER BY cust_id DESC
	LIMIT 1`

	insertCustomer = "InsertCustomer"
	qInsertCustomer = `
	INSERT INTO t_customer
		(cust_id,
		cust_name,
		cust_username,
		cust_password,
		cust_phone,
		cust_email,
		cust_address)
	VALUES(?, ?, ?, ?, ?,  ?, ?)`

	getCustByLogin = "GetCustByLogin"
	qGetCustByLogin = `
	SELECT cust_id,
		cust_name,
		cust_username,
		cust_password,
		cust_phone,
		cust_email,
		cust_address 
	FROM t_customer
	WHERE  
	cust_username = ?
	AND cust_password = ?`

	updateCustomerById = "UpdateCustomerById"
	qUpdateCustomerById = `
	UPDATE t_customer
	SET cust_name = ?,
		cust_username = ?,
		cust_password = ?,
		cust_phone = ?,
		cust_email = ?,
		cust_address = ?
	WHERE cust_id= ?`

	////__________________________________________ T_Product ____________________________________________
	//  - belum -
	getAllProduct  = "GetAllProduct"
	qGetAllProduct = `
	SELECT prod_id,
		ctg_id,
		prod_name,
		prod_stock,
		prod_price,
		prod_desc
	FROM t_product`
	// VALUES (?, ?, ?, ?, ?, ?)`

	insertProduct  = "InsertProduct"
	qInsertProduct = `
	INSERT INTO t_product
		(prod_id,
		ctg_id,
		prod_name,
		prod_stock,
		prod_price,
		prod_desc)
	VALUES (?, ?, ?, ?, ?, ?)`

////__________________________________________ T_Category____________________________________________

	getAllCategory = "GetAllCategory"
	qGetAllCategory = `
	SELECT ctg_id, ctg_type FROM t_category`

////__________________________________________ TH_Order ____________________________________________

	getAllOrder = "GetAllOrder" 
	qGetAllOrder = `
	SELECT ord_id,
		customer_id, 
		ord_status,
		ord_total,
		ord_payment	
	FROM th_order`		

////__________________________________________ TD_Order ____________________________________________

	getAllOrderDetail = "GetAllOrderDetail" 
	qGetAllOrderDetail = `
	SELECT orddetail_id,
		order_id, 
		product_id,
		orddetail_qty,
		orddetail_total		
	FROM td_order`						

////__________________________________________ T_Delivery ____________________________________________


///___________________________________________ JOIN TABLES = T_Admin & T_Customer ____________________________________
	
	// belum tentu bener
	getJoinAdmCust = "GetJoinAdmCust"
	qGetJoinAdmCust =`
	SELECT a.adm_id, a.adm_username, a.adm_password,
       c.cust_id, c.cust_username, c.cust_password
  	FROM t_admin a, t_customer c`

)

var (
	readStmt = []statement{
		//jadwalsttk
		{getAllAdmin, qGetAllAdmin},
		{getAllProduct, qGetAllProduct},
		{getAllCategory, qGetAllCategory},
		{getAllOrder,qGetAllOrder},
		{getAllOrderDetail, qGetAllOrderDetail},
		{getCustById, qGetCustById},
		{getCustByLogin, qGetCustByLogin},
		{getAdmByLogin, qGetAdmByLogin},
		{getCountCust, qGetCountCust},
		{getAdmLastData, qGetAdmLastData},
		{getCustLastData,qGetCustLastData},
		

		{getJoinAdmCust, qGetJoinAdmCust},
	}
	insertStmt = []statement{
		{insertProduct,qInsertProduct},
		{insertCustomer, qInsertCustomer},
		{insertAdmin, qInsertAdmin},
	}
	updateStmt = []statement{
		{updateCustomerById, qUpdateCustomerById},
	}
	deleteStmt = []statement{}
)

// New ...
func New(db *sqlx.DB, tracer opentracing.Tracer, logger jaegerLog.Factory) *Data {
	var (
		stmts = make(map[string]*sqlx.Stmt)
	)
	d := &Data{
		db:     db,
		tracer: tracer,
		logger: logger,
		stmt:   &stmts,
	}

	d.InitStmt()
	return d
}

func (d *Data) InitStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize select statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range insertStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize insert statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range updateStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize update statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range deleteStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize delete statement key %v, err : %v", v.key, err)
		}
	}

	*d.stmt = stmts
}

// contoh implementasi ...
// func (d Data) GetShowname(ctx context.Context, movieID string) (string, error) {
// 	var (
// 		showname string
// 		err      error
// 	)

//// WAJIB ADA
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span := d.tracer.StartSpan("SQL SELECT", opentracing.ChildOf(span.Context()))
// 		span.SetTag("mysql.server", "123.72.156.4")
// 		span.SetTag("mysql.database", "movie")
// 		span.SetTag("mysql.table", "showname")
// 		span.SetTag("mysql.query", "SELECT * FROM movie.showname WHERE movie_id="+movieID)
// 		defer span.Finish()
// 		ctx = opentracing.ContextWithSpan(ctx, span)
// 	}
//// WAJIB ADA

// 	// assumed data fetched from database
// 	showname = "Joni Bizarre Adventure"

//// OPTIONAL, DISARANKAN DIBUAT LOGGINGNYA
// 	d.logger.For(ctx).Info("SQL Query Success", zap.String("showname", showname))

//// WAJIB ADA, INI MERUPAKAN LOGGING KALAU TERJADI ERROR, BISA DIPASANG DI SERVICE DAN HANDLER, TIDAK HANYA DI DATA SAJA
// 	// if err != nil {
// 	// 	d.logger.For(ctx).Error("SQL Query Failed", zap.Error(err))
// 	// 	return showname, err
// 	// }
//// WAJIB ADA

// 	return showname, err
// }
