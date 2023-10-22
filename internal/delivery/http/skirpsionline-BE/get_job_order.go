package skirpsionlineBE

import (
	"log"
	"net/http"
	httpHelper "skripsi-online-BE/internal/delivery/http"
	"skripsi-online-BE/pkg/response"

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

	case "tokenuser":
		err = h.skripsionlineSvc.TokenUser(ctx)

	case "getcustbylogin":
		result, err = h.skripsionlineSvc.GetCustByLogin(ctx, r.FormValue("username"), r.FormValue("password"))
		log.Println("getcustbylogin", r.FormValue("username"), r.FormValue("password"))

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

	case "getjoinadmcust":
		result, err = h.skripsionlineSvc.GetJoinAdmCust(ctx)

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
