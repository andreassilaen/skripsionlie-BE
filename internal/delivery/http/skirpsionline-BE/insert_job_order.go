package skirpsionlineBE

import (
	"skripsi-online-BE/pkg/response"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"

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
		log.Println("header Delivery : ", header)

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
