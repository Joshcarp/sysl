// Code generated by sysl DO NOT EDIT.
package simple

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.service.anz/sysl/server-lib/common"
	"github.service.anz/sysl/server-lib/core"
	"github.service.anz/sysl/server-lib/handlerinitialiser"
	"github.service.anz/sysl/server-lib/validator"
)

// GenCallback callbacks used by the generated code
type GenCallback interface {
	AddMiddleware(ctx context.Context, r chi.Router)
	BasePath() string
	Config() validator.Validator
	HandleError(ctx context.Context, w http.ResponseWriter, kind common.Kind, message string, cause error)
	DownstreamTimeoutContext(ctx context.Context) (context.Context, context.CancelFunc)
}

// Router interface for Simple
type Router interface {
	Route(router *chi.Mux)
}

// ServiceRouter for Simple API
type ServiceRouter struct {
	gc               GenCallback
	svcHandler       *ServiceHandler
	basePathFromSpec string
}

// NewServiceRouter creates a new service router for Simple
func NewServiceRouter(gc GenCallback, svcHandler *ServiceHandler) handlerinitialiser.HandlerInitialiser {
	return &ServiceRouter{gc, svcHandler, "simple"}
}

// WireRoutes ...
//nolint:funlen
func (s *ServiceRouter) WireRoutes(ctx context.Context, r chi.Router) {
	r.Route(core.SelectBasePath(s.basePathFromSpec, s.gc.BasePath()), func(r chi.Router) {
		s.gc.AddMiddleware(ctx, r)
		r.Get("/stuff", s.svcHandler.GetStuffListHandler)
	})
}

// Config ...
func (s *ServiceRouter) Config() validator.Validator {
	return s.gc.Config()
}

// Name ...
func (s *ServiceRouter) Name() string {
	return "Simple"
}
