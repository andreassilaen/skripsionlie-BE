package joborder

import (

	// apichcEntity "sttk/internal/entity/apichc"

	// sttk "sttk/internal/"
	jaegerLog "job-order-be/pkg/log"

	"github.com/opentracing/opentracing-go"
)

// IsttkSvc is an interface to sttk Service
type IjoborderSvc interface {
}

type (
	// Handler ...
	Handler struct {
		joborderSvc IjoborderSvc
		tracer  opentracing.Tracer
		logger  jaegerLog.Factory
	}
)

// New for bridging product handler initialization
func New(is IjoborderSvc, tracer opentracing.Tracer, logger jaegerLog.Factory) *Handler {
	return &Handler{
		joborderSvc: is,
		tracer:  tracer,
		logger:  logger,
	}
}
