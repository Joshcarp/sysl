// Code generated by sysl DO NOT EDIT.
package simple

import (
	"net/http"

	"github.service.anz/sysl/server-lib/common"
	"github.service.anz/sysl/server-lib/restlib"
	"github.service.anz/sysl/server-lib/validator"
)

// Handler interface for simple
type Handler interface {
	GetFoobarListHandler(w http.ResponseWriter, r *http.Request)
	GetTestListHandler(w http.ResponseWriter, r *http.Request)
}

// ServiceHandler for simple API
type ServiceHandler struct {
	genCallback      GenCallback
	serviceInterface *ServiceInterface
}

// NewServiceHandler for simple
func NewServiceHandler(genCallback GenCallback, serviceInterface *ServiceInterface) *ServiceHandler {
	return &ServiceHandler{genCallback, serviceInterface}
}

// GetFoobarListHandler ...
func (s *ServiceHandler) GetFoobarListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetFoobarList == nil {
		s.genCallback.HandleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetFoobarListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	client := GetFoobarListClient{}

	valErr := validator.Validate(&req)
	if valErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	foo, err := s.serviceInterface.GetFoobarList(ctx, &req, client)
	if err != nil {
		s.genCallback.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, foo, err)
}

// GetTestListHandler ...
func (s *ServiceHandler) GetTestListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetTestList == nil {
		s.genCallback.HandleError(r.Context(), w, common.InternalError, "not implemented", nil)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetTestListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	client := GetTestListClient{}

	valErr := validator.Validate(&req)
	if valErr != nil {
		s.genCallback.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr)
		return
	}

	stuff, err := s.serviceInterface.GetTestList(ctx, &req, client)
	if err != nil {
		s.genCallback.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	headermap.Add("Content-Type", "application/json")
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, stuff, err)
}
