package http

import (
	"net/http"

	"job-order-be/pkg/grace"

	"github.com/rs/cors"
)

// SttkHandler ...
type SttkHandler interface {
	GetJobOrder(w http.ResponseWriter, r *http.Request)
	InsertJobOrder(w http.ResponseWriter, r *http.Request)
	DeleteJobOrder(w http.ResponseWriter, r *http.Request)
	UpdateJobOrder(w http.ResponseWriter, r *http.Request)
	// PrintSelisih(w http.ResponseWriter, r *http.Request)
	// PrintExpiredTerpajang(w http.ResponseWriter, r *http.Request)
	// PrintExpiredTerkumpul(w http.ResponseWriter, r *http.Request)

	// PrintBatch(w http.ResponseWriter, r *http.Request)
	// PrintBatchFull(w http.ResponseWriter, r *http.Request)

	// //Trans Out
	// InsertTransOut(w http.ResponseWriter, r *http.Request)
	// InsertSales(w http.ResponseWriter, r *http.Request)
	// DeleteSalesByPeriod(w http.ResponseWriter, r *http.Request)
	// RemoveSalesByOutcode(w http.ResponseWriter, r *http.Request)
	// InsertBatchData(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	Joborder SttkHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
