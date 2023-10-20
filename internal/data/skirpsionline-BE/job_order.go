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
		(adm_name,
		adm_username,
		adm_password,
		adm_phone,
		adm_email,
		adm_address)
	VALUES(?, ?, ?, ?,  ?, ?)`



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


	////__________________________________________ T_Employee____________________________________________
	//	+ selesai +
	getAllEmployee  = "GetAllEmployee"
	qGetAllEmployee = `
	SELECT emp_id,
		emp_name,
		emp_username,
		emp_password,
		emp_phone,
		emp_email,
		emp_address 				
	FROM t_employee`

	getEmpLastData = "GetEmpLastData"
	qGetEmpLastData = `
	SELECT * FROM t_employee
	ORDER BY emp_id DESC
	LIMIT 1`

	insertEmployee = "InsertEmployee"
	qInsertEmployee = `
	INSERT INTO t_employee
		(emp_name,
		emp_username,
		emp_password,
		emp_phone,
		emp_email,
		emp_address )
	VALUES(?, ?, ?,?, ?, ?)`



	getEmpByLogin = "GetEmpByLogin"
	qGetEmpByLogin = `
	SELECT emp_id,
		emp_name,
		emp_username,
		emp_password,
		emp_phone,
		emp_email,
		emp_address 								
	FROM t_employee	
	WHERE  emp_username = ?
	AND emp_password = ?`



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
		(cust_name,
		cust_username,
		cust_password,
		cust_phone,
		cust_email,
		cust_address)
	VALUES( ?, ?, ?, ?,  ?, ?)`

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
		adm_id,
		ctg_id,
		prod_name,
		prod_desc,
		prod_price,
		prod_stock,
		prod_lastupdate,
		prod_img
	FROM t_product`
	// VALUES (?, ?, ?, ?, ?, ?)`

	insertProduct  = "InsertProduct"
	qInsertProduct = `
	INSERT INTO t_product
		(adm_id,
		ctg_id,
		prod_name,
		prod_desc,
		prod_price,
		prod_stock,
		prod_lastupdate,
		prod_img)
	VALUES(?, ?, ?, ?, ?, ?,  NOW(), ?)`

	getProdLastData = "GetProdLastData"
	qGetProdLastData = `
	SELECT * FROM t_product
	ORDER BY prod_id DESC
	LIMIT 1`

////__________________________________________ T_Category____________________________________________

	getAllCategory = "GetAllCategory"
	qGetAllCategory = `
	SELECT ctg_id, ctg_type FROM t_category`

////__________________________________________ TH_Cart ____________________________________________

	getAllCart = "GetAllCart" 
	qGetAllCart = `
	SELECT 
		cart_id,
		cust_id,
		cart_total, 
		cart_lastupdate
	FROM th_cart`		

	getCartByCustId = "GetCartByCustId"
	qGetCartByCustId = `
	SELECT 
		cart_id,														
		cust_id,
		cart_total, 
		cart_lastupdate
		FROM th_cart
	WHERE cust_id = ?`

	insertHeaderCart = "InsertHeaderCart"
	qInsertHeaderCart = `
	INSERT INTO th_cart (
		cust_id,
		cart_total,
		cart_lastupdate)
		VALUES (?, ?, NOW())`

////__________________________________________ TD_Cart ____________________________________________

	// getAllCartDetail = "GetAllCartDetail" 
	// qGetAllCartDetail = `
	// SELECT 
	// 	cart_id,
	// 	prod_id,
	// 	cartdtl_qty,
	// FROM td_cart`	
	
	insertDetailCart = "InsertDetailCart"
	qInsertDetailCart = `
	INSERT INTO td_cart (
		cart_id,
		prod_id,
		cardtl_qty)
		VALUES (?, ?, ?)`
	
	
////__________________________________________ T_Transaction____________________________________________
	
	getTranByCartId = "GetTranByCartId"
	qGetTranByCartId = `
	SELECT 
		tra_id,
		cart_id,
		tra_total,
		tra_img,
		tra_date
	FROM th_transaction
	WHERE cart_id = ?`


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
		{getAllCart,qGetAllCart},
		// {getAllCartDetail, qGetAllCartDetail},
		{getCustById, qGetCustById},
		{getCustByLogin, qGetCustByLogin},
		{getAdmByLogin, qGetAdmByLogin},
		{getCountCust, qGetCountCust},
		{getAdmLastData, qGetAdmLastData},
		{getCustLastData,qGetCustLastData},
		{getProdLastData, qGetProdLastData},

		{getCartByCustId,qGetCartByCustId},

		{getTranByCartId, qGetTranByCartId},

		{getAllEmployee, qGetAllEmployee},	
		{getEmpByLogin, qGetEmpByLogin},	
		{getEmpLastData, qGetEmpLastData},

		// {getJoinAdmCust, qGetJoinAdmCust},
	}
	insertStmt = []statement{
		{insertProduct,qInsertProduct},
		{insertCustomer, qInsertCustomer},
		{insertAdmin, qInsertAdmin},
		{insertEmployee, qInsertEmployee},
		{insertHeaderCart,qInsertHeaderCart},
		{insertDetailCart, qInsertDetailCart},
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
