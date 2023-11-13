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

	getAdmById  = "GetAdmById"
	qGetAdmById = `
	SELECT adm_id,
		adm_name,
		adm_username,
		adm_password,
		adm_phone,
		adm_email,
		adm_address 
	FROM t_admin
	WHERE  
	adm_id = ?`

	getAdmLastData  = "GetAdmLastData"
	qGetAdmLastData = `
	SELECT * FROM t_admin
	ORDER BY adm_id DESC
	LIMIT 1`

	insertAdmin  = "InsertAdmin"
	qInsertAdmin = `
	INSERT INTO t_admin
		(adm_name,
		adm_username,
		adm_password,
		adm_phone,
		adm_email,
		adm_address)
	VALUES(?, ?, ?, ?,  ?, ?)`

	getAdmByLogin  = "GetAdmByLogin"
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

	updateAdminById  = "UpdateAdminById"
	qUpdateAdminById = `
	UPDATE t_admin
	SET adm_name = ?,
		adm_username = ?,
		adm_password = ?,
		adm_phone = ?,
		adm_email = ?,
		adm_address = ?
	WHERE adm_id= ?`

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

	getEmpById  = "GetEmpById"
	qGetEmpById = `
	SELECT emp_id,
		emp_name,
		emp_username,
		emp_password,
		emp_phone,
		emp_email,
		emp_address 
	FROM t_employee
	WHERE  
	emp_id = ?`

	getEmpLastData  = "GetEmpLastData"
	qGetEmpLastData = `
	SELECT * FROM t_employee
	ORDER BY emp_id DESC
	LIMIT 1`

	insertEmployee  = "InsertEmployee"
	qInsertEmployee = `
	INSERT INTO t_employee
		(emp_name,
		emp_username,
		emp_password,
		emp_phone,
		emp_email,
		emp_address )
	VALUES(?, ?, ?,?, ?, ?)`

	updateEmployeeById  = "UpdateEmployeeById"
	qUpdateEmployeeById = `
	UPDATE t_employee
	SET emp_name = ?,
		emp_username = ?,
		emp_password = ?,
		emp_phone = ?,
		emp_email = ?,
		emp_address = ?
	WHERE emp_id= ?`

	getEmpByLogin  = "GetEmpByLogin"
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

	getCustById  = "GetCustById"
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

	getCountCust  = "GetCountCust"
	qGetCountCust = `
	SELECT COUNT(cust_id) AS total
	FROM t_customer`

	getCustLastData  = "GetCustLastData"
	qGetCustLastData = `
	SELECT * FROM t_customer
	ORDER BY cust_id DESC
	LIMIT 1`

	insertCustomer  = "InsertCustomer"
	qInsertCustomer = `
	INSERT INTO t_customer
		(cust_name,
		cust_username,
		cust_password,
		cust_phone,
		cust_email,
		cust_address)
	VALUES( ?, ?, ?, ?,  ?, ?)`

	getCustByLogin  = "GetCustByLogin"
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

	updateCustomerById  = "UpdateCustomerById"
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
	SELECT p.prod_id,
		p.adm_id,
		p.ctg_id,
		c.ctg_type,
		p.prod_name,
		p.prod_desc,
		p.prod_price,
		p.prod_stock,
		p.prod_lastupdate,
		p.prod_img
	FROM t_product p, t_category c
	WHERE p.ctg_id = c.ctg_id`
	// VALUES (?, ?, ?, ?, ?, ?)`

	getProdById  = "GetProdById"
	qGetProdById = `
	SELECT p.prod_id,
		p.adm_id,
		p.ctg_id,
		c.ctg_type,
		p.prod_name,
		p.prod_desc,
		p.prod_price,
		p.prod_stock,
		p.prod_lastupdate,
		p.prod_img
	FROM t_product p, t_category c
	WHERE p.ctg_id = c.ctg_id
	AND prod_id = ?`

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

	getProdLastData  = "GetProdLastData"
	qGetProdLastData = `
	SELECT * FROM t_product
	ORDER BY prod_id DESC
	LIMIT 1`

	updateProdById  = "UpdateProdById"
	qUpdateProdById = `
	UPDATE t_product
	SET  adm_id = ?,
		prod_name = ?,
		prod_desc = ?,
		prod_price = ?,
		prod_stock = ?,
		prod_lastupdate = NOW(),
		prod_img = ?
	WHERE prod_id= ?`

	deleteProductByProdId = "DeleteProductByProdId"
	qDeleteProductByProdid = `
	DELETE FROM t_product
	WHERE prod_id= ?`


	////__________________________________________ T_Category____________________________________________

	getAllCategory  = "GetAllCategory"
	qGetAllCategory = `
	SELECT ctg_id, ctg_type FROM t_category`

	////__________________________________________ TH_Cart ____________________________________________

	getAllCart  = "GetAllCart"
	qGetAllCart = `
	SELECT 
		cart_id,
		cust_id,
		cart_total, 
		cart_payedyn,
		cart_lastupdate
	FROM th_cart`

	getCartByCustId  = "GetCartByCustId"
	qGetCartByCustId = `
	SELECT 
		cart_id,														
		cust_id,
		cart_total, 
		cart_payedyn,
		cart_lastupdate
		FROM th_cart
	WHERE cust_id = ?`

	getHeaderCartLastData  = "GetHeaderCartLastData"
	qGetHeaderCartLastData = `
	SELECT * FROM th_cart
	ORDER BY cart_id DESC
	LIMIT 1`

	insertHeaderCart  = "InsertHeaderCart"
	qInsertHeaderCart = `
	INSERT INTO th_cart (
		cust_id,
		cart_total,
		cart_lastupdate)
		VALUES (?, ?, NOW())`

	getHeaderCartNotPayedByCustId = "GetHeaderCartNotPayedCustId"
	qGetHeaderCartNotPayedCustId  = `
	SELECT * 
	FROM th_cart 
	WHERE cart_payedyn = "N"
	AND cust_id = ?
	ORDER BY cart_id DESC
	LIMIT 1`

	updateHeaderCartPayed  =  "UpdateHeaderCartPayed"
	qUpdateHeaderCartPayed = `
	UPDATE th_cart
	SET cart_payedyn = "Y",
		cart_lastupdate = NOW()
	WHERE cart_payedyn = "N"
		AND cart_id = ?`

	////__________________________________________ TD_Cart ____________________________________________

	// getAllCartDetail = "GetAllCartDetail"
	// qGetAllCartDetail = `
	// SELECT
	// 	cart_id,
	// 	prod_id,
	// 	cartdtl_qty,
	// FROM td_cart`

	insertDetailCart  = "InsertDetailCart"
	qInsertDetailCart = `
	INSERT INTO td_cart (
		cart_id,
		prod_id,
		cardtl_qty)
		VALUES (?, ?, ?)`

	////__________________________________________ TH_Transaction____________________________________________

	getTranByCartId  = "GetTranByCartId"
	qGetTranByCartId = `
	SELECT 
		tra_id,
		cart_id,
		cust_id,
		rek_id,
		tra_total,
		tra_img,
		tra_date
	FROM th_transaction
	WHERE cart_id = ?`

	getAllHeaderTran  = "GetAllHeaderTran"
	qGetAllHeaderTran = `
	SELECT tra_id,
		cart_id,
		cust_id,
		rek_id,
		tra_total,
		tra_img,
		tra_date
	FROM th_transaction
	WHERE tra_checkedyn = "N"`

	getHeaderTranLastDataByCusId  = "GetHeaderTranLastDataByCusId"
	qGetHeaderTranLastDataByCusId = `
	SELECT * FROM th_transaction
	WHERE cust_id = ?
	ORDER BY tra_id DESC
	LIMIT 1`

	insertHeaderTran  = "InsertHeaderTran"
	qInsertHeaderTran = `
	INSERT INTO th_transaction(
		cart_id,
		cust_id,
		rek_id,
		tra_total,
		tra_img,
		tra_date)
	VALUES (?, ?, ?, ?, ?, NOW())`


	getAllHeaderTranByCustId = "GetAllHeaderTranByCustId"
	qGetAllHeaderTranByCustId = `
	SELECT * 
	FROM th_transaction 
	WHERE cust_id = ?`

	updateTHTranChecked  = "UpdateTHTranChecked"
	qUpdateTHTranChecked = `
	UPDATE th_transaction 
	SET tra_checkedyn = "Y",
		tra_date = now()
	WHERE tra_id = ?`
	

	////__________________________________________ TD_Transaction____________________________________________

	getDetailTranByTraId  = "GetDetailTranByTraId"
	qGetDetailTranByTraId = `
	SELECT 
		tra_id,
		prod_id,
		tradtl_qty,
		tradtl_price,
		tradtl_amount
	FROM td_transaction
	WHERE tra_id = ?`

	insertDetailTran  = "InsertDetailTran"
	qInsertDetailTran = `
	INSERT INTO td_transaction (
		tra_id,
		prod_id,
		tradtl_qty,
		tradtl_price,
		tradtl_amount)
	VALUES (?, ?, ?, ?, ?)`

	////__________________________________________ T_Order ____________________________________________

	getAllOrder  = "GetAllOrder"
	qGetAllOrder = `
	SELECT
		ord_id,
		adm_id, 
		tra_id, 
		ord_confirmedyn,
		ord_lastupdate
	FROM t_order`

	insertOrder  = "InsertOrder"
	qInsertOrder = `
	INSERT INTO t_order (
		adm_id, 
		tra_id, 
		ord_lastupdate)
	VALUES (?, ?, NOW())`

	insertOrderAcc  = "InsertOrderAcc"
	qInsertOrderAcc = `
	INSERT INTO t_order (
		adm_id, 
		tra_id, 
		ord_confirmedyn,
		ord_lastupdate)
	VALUES (?, ?, "Y",NOW())`

	updateOrderOnDeliveryYes  = "UpdateOrderOnDeliveryYes"
	qUpdateOrderOnDeliveryYes = `
	UPDATE t_order 
	SET ord_ondeliveryyn = "Y",
		ord_lastupdate = now()
	WHERE ord_id = ?`

	////__________________________________________ T_Delivery ____________________________________________

	getAllDelivery  = "GetAllDelivery"
	qGetAllDelivery = `
	SELECT * 
	FROM t_delivery`

	getDeliveryByEmpId  = "GetDeliverByEmpId"
	qGetDeliveryByEmpId = `
	SELECT * 
	FROM t_delivery
	WHERE emp_id = ?
	ORDER BY delivery_doneyn ASC,
	  delivery_date ASC`

	insertDeliveryProcess  = "InsertDeliveryProcess"
	qInsertDeliveryProcess = `
	INSERT INTO t_delivery (
		emp_id, 
		ord_id,
		delivery_doneyn,
		delivery_date)
	VALUES (?, ?, "N", NOW())`

	updateDeliveryDone = "UpdateDeliveryDone"
	qUpdateDeliveryDone = `
	UPDATE t_delivery
	SET delivery_doneyn = "Y",
		delivery_date= NOW()
	WHERE ord_id = ?`

	////__________________________________________ T_Rekening ____________________________________________

	getAllRekening  = "GetAllRekening"
	qGetAllRekening = `
	SELECT * FROM t_rekening`

	getRekByRekId  = "GetRekByRekId"
	qGetRekByRekId = `
	SELECT * 
	FROM t_rekening
	WHERE rek_id = ? `

	///___________________________________________ JOIN TABLES ____________________________________

	// belum tentu bener
	getJoinAdmCust  = "GetJoinAdmCust"
	qGetJoinAdmCust = `
	SELECT a.adm_id, a.adm_username, a.adm_password,
       c.cust_id, c.cust_username, c.cust_password
  	FROM t_admin a, t_customer c`

	getJoinOrdCustTHTra  = "GetJoinOrdCustTHTra"
	qGetJoinOrdCustTHTra = `
	SELECT 
		o.ord_id,
		o.tra_id, 
		o.adm_id, 
		c.cust_name, 
		c.cust_address, 
		h.tra_total, 
		o.ord_lastupdate
	FROM t_order o, t_customer c, th_transaction h
	WHERE 
		o.tra_id = h.tra_id
		AND h.cust_id = c.cust_id
		AND o.ord_confirmedyn = "Y"
		and o.ord_ondeliveryyn = "N"`

	getJoinOrdCustTHTraByOrdId  = "GetJoinOrdCustTHTraByOrdId"
	qGetJoinOrdCustTHTraByOrdId = `
	SELECT 
		o.ord_id,
		o.tra_id,
		o.adm_id, 
		c.cust_name, 
		c.cust_address, 
		h.tra_total, 
		o.ord_lastupdate
	FROM t_order o, t_customer c, th_transaction h
	WHERE 
		o.tra_id = h.tra_id
		AND h.cust_id = c.cust_id 
		AND o.ord_id = ?`

	getJoinTDTraProdByTraId  = "GetJoinTDTraProdByTraId"
	qGetJoinTDTraProdByTraId = `
	SELECT 
		d.tra_id,  
		p.prod_name, 
		d.tradtl_qty, 
		d.tradtl_amount,
		c.cust_name,
		c.cust_address
	FROM th_transaction h, td_transaction d, t_product p, t_customer c
	WHERE h.tra_id = d.tra_id
		AND d.prod_id = p.prod_id
		AND h.cust_id = c.cust_id
		AND d.tra_id = ?`

	// old query dari get fun diatatas
	// 	SELECT
	// 	d.tra_id,
	// 	p.prod_name,
	// 	d.tradtl_qty,
	// 	d.tradtl_amount
	// FROM td_transaction d, t_product p
	// WHERE
	// 	d.prod_id = p.prod_id
	// 	AND d.tra_id = ?

	getJoinOrdTHTDTraProdByOrdId = "GetJoinOrdTHTDTraProdByOrdId"
	qGetJoinOrdTHTDTraProdByOrdId = `
	SELECT 
		o.ord_id,
		d.tra_id,
		p.prod_name, 
		d.tradtl_qty, 
		d.tradtl_amount,
		c.cust_name,
		c.cust_address
	FROM t_delivery del, t_order o, th_transaction h, td_transaction d, t_product p, t_customer c
	WHERE o.ord_id = del.ord_id
		AND o.tra_id = h.tra_id
		AND h.tra_id = d.tra_id
		AND d.prod_id = p.prod_id
		AND h.cust_id = c.cust_id

		AND del.ord_id= ?`
	// AND delivery_doneyn = "N"

	getListJoinTHTDCartProdByCustIdAndCartId  = "GetListJoinTHTDCartProdByCustIdAndCartId"
	qGetListJoinTHTDCartProdByCustIdAndCartId = `
	SELECT h.cart_id, h.cust_id, h.cart_total, h.cart_payedyn, p.prod_id, p.prod_name, p.prod_desc, p.prod_price, p.prod_stock, d.cardtl_qty, p.prod_img
	FROM th_cart h, td_cart d, t_product p 
	WHERE h.cart_id = d.cart_id
	AND d.prod_id = p.prod_id
	AND h.cart_payedyn = "N"
	AND h.cust_id = ?
	AND d.cart_id = ? `

	getProductInJoinTHTDCartProdByProdId  = "GetProductInJOinTHTDCartProdByProdId"
	qGetProductInJOinTHTDCartProdByProdId = `
	SELECT h.cart_id, h.cust_id, h.cart_total, h.cart_payedyn, p.prod_id, p.prod_name, p.prod_desc, p.prod_price, p.prod_stock, d.cardtl_qty, p.prod_img
	FROM th_cart h, td_cart d, t_product p 
	WHERE h.cart_id = d.cart_id
	AND d.prod_id = p.prod_id
	AND h.cart_payedyn = "N"
	AND h.cust_id = ?
	AND d.cart_id = ?
	AND d.prod_id = ? `

	updateQtyDetailJoinTHTDCart  = "UpdateQtyDetailJoinTHTDCart"
	qUpdateQtyDetailJoinTHTDCart = `
	UPDATE th_cart h, td_cart d
	SET 
		h.cart_lastupdate = NOW(),
		d.cardtl_qty = ?
	WHERE h.cart_id = d.cart_id
	AND h.cust_id = ?
	AND d.cart_id = ?
	AND prod_id = ?`

	getJoinTHTraRekByCusId  = "GetJoinTHTraRekByCusId"
	qGetJoinTHTraRekByCusId = `
	SELECT h.tra_id, h.cust_id, r.rek_bank, h.tra_total, h.tra_img, h.tra_date, h.tra_checkedyn
	FROM th_transaction h, t_rekening r
	WHERE h.rek_id = r.rek_id
	AND h.cust_id = ? `

	getJoinOrdTHTraByCustId = "GetJoinOrdTHTraByCustId"
	qGetJoinOrdTHTraByCustId = `
	SELECT 
		o.ord_id,  
		h.tra_id, 
		h.tra_total,  
		o.ord_confirmedyn,  
		o.ord_ondeliveryyn, 
		o.ord_lastupdate,
		d.delivery_doneyn
	FROM t_order o, th_transaction h, t_delivery d
	WHERE o.tra_id = h.tra_id 
		AND o.ord_id = d.ord_id
		AND h.cust_id = ?`


	getCountDashboardAdmin = "GetCountDashboardAdmin"
	qGetCountDashboardAdmin = `
	SELECT 
	(SELECT COUNT(*) FROM th_transaction h WHERE h.tra_checkedyn = "N") AS tra_uncheck,
	(SELECT COUNT(*) FROM t_order WHERE ord_confirmedyn = 'N' AND ord_ondeliveryyn = 'N') AS ord_canceled,
	(SELECT COUNT(*) FROM t_order WHERE ord_confirmedyn = 'Y' AND ord_ondeliveryyn = 'N') AS ord_process,
	(SELECT COUNT(*) FROM t_order WHERE ord_confirmedyn = 'Y' AND ord_ondeliveryyn = 'Y') AS ord_delivery,
	(SELECT COUNT(*) FROM t_delivery d WHERE d.delivery_doneyn = 'N') AS del_ongoing ,
	(SELECT COUNT(*) FROM t_delivery d WHERE d.delivery_doneyn = 'Y') AS del_doned  `

)

var (
	readStmt = []statement{
		//jadwalsttk
		{getAllAdmin, qGetAllAdmin},
		{getAllProduct, qGetAllProduct},
		{getAllCategory, qGetAllCategory},
		{getAllCart, qGetAllCart},
		// {getAllCartDetail, qGetAllCartDetail},
		{getCustById, qGetCustById},
		{getAdmById, qGetAdmById},
		{getEmpById, qGetEmpById},
		{getProdById, qGetProdById},
		{getCustByLogin, qGetCustByLogin},
		{getAdmByLogin, qGetAdmByLogin},
		{getCountCust, qGetCountCust},
		{getAdmLastData, qGetAdmLastData},
		{getCustLastData, qGetCustLastData},
		{getProdLastData, qGetProdLastData},
		{getHeaderCartLastData, qGetHeaderCartLastData},

		{getCartByCustId, qGetCartByCustId},
		{getHeaderCartNotPayedByCustId, qGetHeaderCartNotPayedCustId},

		{getAllHeaderTran, qGetAllHeaderTran},
		{getTranByCartId, qGetTranByCartId},
		{getHeaderTranLastDataByCusId, qGetHeaderTranLastDataByCusId},
		{getAllHeaderTranByCustId, qGetAllHeaderTranByCustId},

		{getDetailTranByTraId, qGetDetailTranByTraId},

		{getAllEmployee, qGetAllEmployee},
		{getEmpByLogin, qGetEmpByLogin},
		{getEmpLastData, qGetEmpLastData},

		{getAllOrder, qGetAllOrder},
		{getDeliveryByEmpId, qGetDeliveryByEmpId},

		{getAllDelivery, qGetAllDelivery},

		{getAllRekening, qGetAllRekening},
		{getRekByRekId, qGetRekByRekId},

		// {getJoinAdmCust, qGetJoinAdmCust},
		{getJoinOrdCustTHTra, qGetJoinOrdCustTHTra},
		{getJoinOrdCustTHTraByOrdId, qGetJoinOrdCustTHTraByOrdId},
		{getJoinTDTraProdByTraId, qGetJoinTDTraProdByTraId},
		{getJoinOrdTHTDTraProdByOrdId, qGetJoinOrdTHTDTraProdByOrdId},
		{getProductInJoinTHTDCartProdByProdId, qGetProductInJOinTHTDCartProdByProdId},
		{getListJoinTHTDCartProdByCustIdAndCartId, qGetListJoinTHTDCartProdByCustIdAndCartId},
		{getJoinTHTraRekByCusId, qGetJoinTHTraRekByCusId},
		{getJoinOrdTHTraByCustId, qGetJoinOrdTHTraByCustId},
		{getCountDashboardAdmin, qGetCountDashboardAdmin},
	}
	insertStmt = []statement{
		{insertProduct, qInsertProduct},
		{insertCustomer, qInsertCustomer},
		{insertAdmin, qInsertAdmin},
		{insertEmployee, qInsertEmployee},
		{insertHeaderCart, qInsertHeaderCart},
		{insertDetailCart, qInsertDetailCart},
		{insertHeaderTran, qInsertHeaderTran},
		{insertDetailTran, qInsertDetailTran},
		{insertOrder, qInsertOrder},
		{insertOrderAcc, qInsertOrderAcc},
		{insertDeliveryProcess, qInsertDeliveryProcess},
	}
	updateStmt = []statement{
		{updateAdminById, qUpdateAdminById},
		{updateEmployeeById, qUpdateEmployeeById},
		{updateCustomerById, qUpdateCustomerById},
		{updateProdById, qUpdateProdById},
		{updateQtyDetailJoinTHTDCart, qUpdateQtyDetailJoinTHTDCart},
		{updateOrderOnDeliveryYes, qUpdateOrderOnDeliveryYes},
		{updateDeliveryDone, qUpdateDeliveryDone},
		{updateHeaderCartPayed, qUpdateHeaderCartPayed},
		{updateTHTranChecked, qUpdateTHTranChecked},
	}
	deleteStmt = []statement{
		{deleteProductByProdId, qDeleteProductByProdid},
	}
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
