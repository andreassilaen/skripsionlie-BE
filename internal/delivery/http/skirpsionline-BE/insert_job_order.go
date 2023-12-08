package skirpsionlineBE

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
func (h *Handler) InsertSkripsiOnlineBE(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error

		resp  response.Response
		types string

		// traid int
		reknumber int

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
	case "":

	case "insertproduct":
		var header SBeEntity.InsertProduct
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertProduct(ctx, header)
		log.Println("Delivery InsertProduct : ", header)

	case "insertcustomer":
		var header SBeEntity.InsertCustomer
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertCustomer(ctx, header)
		log.Println("Delivery InsertCustomer : ", header)

	case "insertadmin":
		var header SBeEntity.InsertAdmin
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertAdmin(ctx, header)
		log.Println("Delivery InsertAdmin : ", header)

	case "insertemployee":
		var header SBeEntity.InsertEmployee
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertEmployee(ctx, header)
		log.Println("Delivery InsertEmployee : ", header)

	case "insertheadercart":
		var header SBeEntity.InsertHeaderCart
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertHeaderCart(ctx, header)
		log.Println("Delivery InsertHeaderCart : ", header)

	case "insertdetailcart":
		var header SBeEntity.InsertDetailCart
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertDetailCart(ctx, header)
		log.Println("Delivery InsertDetailCart : ", header)

	case "insertheadertran":
		var header SBeEntity.InsertHeaderTransaction
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertHeaderTran(ctx, header)
		log.Println("Delivery InsertHeaderTran : ", header)

	case "insertdetailtran":
		var header SBeEntity.InsertDetailTransaction
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertDetailTransaction(ctx, header)
		log.Println("Delivery InsertDetailTransaction : ", header)

	case "insertorder":
		var header SBeEntity.InsertOrder
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertOrder(ctx, header)
		log.Println("Delivery InsertDetailCart : ", header)

	case "insertdeliveryprocess":
		var header SBeEntity.InsertDelivery
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertDeliveryProcess(ctx, header)
		log.Println("Delivery InsertDeliveryProcess : ", header)

	case "insertorderacc":
		var header SBeEntity.InsertOrder
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertOrderAcc(ctx, header)
		log.Println("Delivery InsertOrderAcc : ", header)

	case "insertorderbyadmincancelmain":
		// traid, _ = strconv.Atoi(r.FormValue("traid"))
		var header SBeEntity.InsertOrder
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertOrderByAdminCancelMain(ctx, header)
		log.Println("Delivery InsertOrderByAdminMain : ", header)

	case "insertorderbyadminaccmain":
		// traid, _ = strconv.Atoi(r.FormValue("traid"))
		var header SBeEntity.InsertOrder
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertOrderByAdminAccMain(ctx, header)
		log.Println("Delivery InsertOrderByAdminAccMain : ", header)

	case "checkuser":
		var header SBeEntity.CheckUser
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, metadata, err = h.skripsionlineSvc.CheckUser(ctx, header.CheckUserBody)
		log.Println("Delivery CheckUser : ", header)

	case "insertjoinheaderdetailcart":
		var header SBeEntity.InsertJoinHeaderDetailCart
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertJoinHeaderDetailCart(ctx, header)
		log.Println("Delivery InsertJoinHeaderDetailCart : ", header)

	case "insertjoinheaderdetailtran":
		var header SBeEntity.InsertJoinHeaderDetailTran
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &header)
		result, err = h.skripsionlineSvc.InsertJoinHeaderDetailTran(ctx, header)
		log.Println("Delivery InsertJoinHeaderDetailTran : ", header)


	case "insertcategory":
		result, err = h.skripsionlineSvc.InsertCategory(ctx, r.FormValue("ctgtype"))
		log.Println("Delivery insertcategory : ", r.FormValue("ctgtype"))

	case "insertrekening":
		reknumber, _ = strconv.Atoi(r.FormValue("reknumber"))
		result, err = h.skripsionlineSvc.InsertRekening(ctx, r.FormValue("rekbank"), reknumber, r.FormValue("rekname"))
		log.Println("Delivery InsertRekening : ", r.FormValue("rekbank"), reknumber, r.FormValue("rekname"))


	}

	if err != nil {
		resp.SetError(err, http.StatusInternalServerError)
		resp.StatusCode = 500
		log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	return
}

// func (h *Handler) InsertTransOut(w http.ResponseWriter, r *http.Request) {
// 	var (
// 		resp    response.Response
// 		err     error
// 		payload []sttkEntity.TransOut
// 	)

// 	ctx := r.Context()
// 	defer resp.RenderJSON(w, r)

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		resp.SetError(err, http.StatusInternalServerError)
// 		resp.StatusCode = 500
// 		log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
// 		return
// 	}
// 	err = json.Unmarshal(body, &payload)
// 	if err != nil {
// 		resp.SetError(err, http.StatusInternalServerError)
// 		resp.StatusCode = 500
// 		log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
// 		return
// 	}

// 	err = h.sttkSvc.InsertTransOut(ctx, payload)
// 	if err != nil {
// 		resp.SetError(err, http.StatusInternalServerError)
// 		log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
// 		return
// 	}

// 	resp.Error = response.Error{
// 		Status: false,
// 		Msg:    "Berhasil",
// 		Code:   200,
// 	}

// 	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
// }
