package skirpsionlineBE

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	httpHelper "skripsi-online-BE/internal/delivery/http"
	"skripsi-online-BE/pkg/response"
	"strconv"

	SBeEntity "skripsi-online-BE/internal/entity/skirpsionline-BE"

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
func (h *Handler) UpdateSkripsiOnlineBE(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
		types    string

		ordid     int
		tradid    int
		cartid    int
		prodid    int
		prodstock int
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
	case "updateadminbyid":
		var header SBeEntity.UpdateAdminById
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.UpdateAdminById(ctx, header, r.FormValue("admid"))
		log.Println("UpdateAdminById", header, r.FormValue("admid"))

	case "updateemployeebyid":
		var header SBeEntity.UpdateEmployeeById
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.UpdateEmployeeById(ctx, header, r.FormValue("empid"))
		log.Println("UpdateEmployeeById", header, r.FormValue("empid"))

	case "updatethtranchecked":
		tradid, _ = strconv.Atoi(r.FormValue("tradid"))
		result, err = h.skripsionlineSvc.UpdateTHTranChecked(ctx, tradid)
		log.Println("UpdateTHTranChecked", tradid)

	case "updateusermain":
		var header SBeEntity.UpdateUserMain
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.UpdateUserMain(ctx, header, r.FormValue("userid"), r.FormValue("role"))
		log.Println("updateusermain => ", r.FormValue("role"), r.FormValue("userid"), header)

	case "updatecustomerbyid":
		var header SBeEntity.UpdateCustomerById
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.UpdateCustomerById(ctx, header, r.FormValue("cusid"))
		log.Println("UpdateCustomerById", header, r.FormValue("cusid"))

	case "updateprodbyid":
		var header SBeEntity.UpdateProdById
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.UpdateProdById(ctx, header, r.FormValue("prodid"))
		log.Println("UpdateProdById", header, r.FormValue("prodid"))

	case "updateqtydetailjointhtdcart":
		var header SBeEntity.UpdateQtyDetailJoinTHTDCartProd
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.UpdateQtyDetailJoinTHTDCart(ctx, header)
		log.Println("UpdateQtyDetailJoinTHTDCartProd", header)
		// case "insertcustomer":
		// 	var header SBeEntity.InsertCustomer
		// 	body, _ := ioutil.ReadAll(r.Body)
		// 	json.Unmarshal(body, &header)
		// 	result, err = h.skripsionlineSvc.InsertCustomer(ctx, header)
		// 	log.Println("Delivery InsertCustomer : ", header)

	case "updateorderondeliveryyes":
		ordid, _ = strconv.Atoi(r.FormValue("ordid"))
		result, err = h.skripsionlineSvc.UpdateOrderOnDeliveryYes(ctx, ordid)
		log.Println("UpdateOrderOnDeliveryYes", r.FormValue("ordid"))

	case "updatedeliverydone":
		var header SBeEntity.UpdateDeliveryImg
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		// ordid, _ = strconv.Atoi(r.FormValue("ordid"))
		result, err = h.skripsionlineSvc.UpdateDeliveryDone(ctx, r.FormValue("ordid"), header)
		log.Println("UpdateDeliveryDone +>>>", r.FormValue("ordid"), header)

	case "updateheadercartpayed":
		cartid, _ = strconv.Atoi(r.FormValue("cartid"))
		result, err = h.skripsionlineSvc.UpdateHeaderCartPayed(ctx, cartid)
		log.Println("Delivery UpdateHeaderCartPayed : ", ordid)

	case "updateprodstockbyid":
		prodid, _ = strconv.Atoi(r.FormValue("prodid"))
		prodstock, _ = strconv.Atoi(r.FormValue("prodstock"))
		result, err = h.skripsionlineSvc.UpdateProdStockById(ctx, prodstock, prodid)
		log.Println("UpdateProdStockById =>", prodstock, prodid)

	case "updatestock":
		tradid, _ = strconv.Atoi(r.FormValue("tradid"))
		prodstock, _ = strconv.Atoi(r.FormValue("prodstock"))
		result, err = h.skripsionlineSvc.UpdateStock(ctx, (r.FormValue("tradid")))
		log.Println("UpdateStock =>", (r.FormValue("tradid")))

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
