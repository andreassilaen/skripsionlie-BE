package skirpsionlineBE

import (
	"log"
	"net/http"
	httpHelper "skripsi-online-BE/internal/delivery/http"
	"skripsi-online-BE/pkg/response"
	"strconv"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
)

// Getskripsionline godoc
// @Summary Get entries of all sttks
// @Description Get entries of all sttks
// @Tags sttk
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200

// @Router /v1/profiles [get]
func (h *Handler) GetSkripsiOnlineBE(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
		types    string

		ordid  int
		rekid	int
	)
	defer resp.RenderJSON(w, r)

	spanCtx, _ := h.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := h.tracer.StartSpan("Getskripsionline", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	ctx := r.Context()
	ctx = opentracing.ContextWithSpan(ctx, span)
	h.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Your code here
	types = r.FormValue("type")
	switch types {
	case "getalladmin":
		result, err = h.skripsionlineSvc.GetAllAdmin(ctx)

	case "getallproduct":
		result, err = h.skripsionlineSvc.GetAllProduct(ctx)

	case "getallcategory":
		result, err = h.skripsionlineSvc.GetAllCategory(ctx)

	case "getallcart":
		result, err = h.skripsionlineSvc.GetAllCart(ctx)

	case "getallheadertran":
		result, err = h.skripsionlineSvc.GetAllHeaderTran(ctx)

	case "getallorder":
		result, err = h.skripsionlineSvc.GetAllOrder(ctx)

	case "getalldelivery":
		result, err = h.skripsionlineSvc.GetAllDelivery(ctx)

	case "getallrekening":
		result, err = h.skripsionlineSvc.GetAllRekening(ctx)

	case "tokenuser":
		err = h.skripsionlineSvc.TokenUser(ctx)

	case "getcustbylogin":
		result, err = h.skripsionlineSvc.GetCustByLogin(ctx, r.FormValue("username"), r.FormValue("password"))
		log.Println("getcustbylogin", r.FormValue("username"), r.FormValue("password"))

	case "getlistjointhtdcartprodbycustidandcartid":
		result, err = h.skripsionlineSvc.GetListJoinTHTDCartProdByCustIdAndCartId(ctx, r.FormValue("custid"), r.FormValue("cartid"))
		log.Println("getlistjointhtdcartprodbycustidandcartid", r.FormValue("custid"), r.FormValue("cartid"))

	case "getcustbyid":
		result, err = h.skripsionlineSvc.GetCustById(ctx, r.FormValue("cusid"))
		log.Println("getcustbyid", r.FormValue("cusid"))

	case "getadmbylogin":
		result, err = h.skripsionlineSvc.GetAdmByLogin(ctx, r.FormValue("username"), r.FormValue("password"))
		log.Println("getadmbylogin", r.FormValue("username"), r.FormValue("password"))

	case "getadmlastdata":
		result, err = h.skripsionlineSvc.GetAdmLastData(ctx)

	case "getcustlastdata":
		result, err = h.skripsionlineSvc.GetCustLastData(ctx)
	case "getprodlastdata":
		result, err = h.skripsionlineSvc.GetProdLastData(ctx)

	case "getallemployee":
		result, err = h.skripsionlineSvc.GetAllEmployee(ctx)
	case "getemplastdata":
		result, err = h.skripsionlineSvc.GetEmpLastData(ctx)
	case "getempbylogin":
		result, err = h.skripsionlineSvc.GetEmpByLogin(ctx, r.FormValue("username"), r.FormValue("password"))
		log.Println("getempbylogin", r.FormValue("username"), r.FormValue("password"))

	case "getcartbycustid":
		result, err = h.skripsionlineSvc.GetCartByCustId(ctx, r.FormValue("cusid"))
		log.Println("getcartbycustid", r.FormValue("cusid"))

	case "gettranbycartid":
		result, err = h.skripsionlineSvc.GetTranByCartId(ctx, r.FormValue("cartid"))
		log.Println("gettranbycartid", r.FormValue("cartid"))

	case "getdetailtranbytraid":
		result, err = h.skripsionlineSvc.GetDetailTranByTraId(ctx, r.FormValue("traid"))
		log.Println("getdetailtranbytraid", r.FormValue("traid"))

	case "getdeliverybyempid":
		result, err = h.skripsionlineSvc.GetDeliverByEmpId(ctx, r.FormValue("empid"))
		log.Println("getdeliverybyempid", r.FormValue("empid"))

	case "getrekbyrekid":
		rekid,_ = strconv.Atoi(r.FormValue("rekid"))
		result, err = h.skripsionlineSvc.GetRekByRekId(ctx, rekid)
		log.Println("getrekbyrekid", rekid)

	case "getjoinadmcust":
		result, err = h.skripsionlineSvc.GetJoinAdmCust(ctx)

	case "getjoinordcustthtra":
		result, err = h.skripsionlineSvc.GetJoinOrdCustTHTra(ctx)

	case "getjoinordcustthtrabyordid":
		ordid,_ = strconv.Atoi(r.FormValue("ordid"))
		result, err = h.skripsionlineSvc.GetJoinOrdCustTHTraByOrdId(ctx, ordid)
		log.Println("getjoinordcustthtrabyordid", ordid)

	case "getjointdtraprodbytraid":
		result, err = h.skripsionlineSvc.GetJoinTDTraProdByTraId(ctx, r.FormValue("traid"))
		log.Println("getjointdtraprodbytraid", r.FormValue("traid"))

	}

	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	return
}
