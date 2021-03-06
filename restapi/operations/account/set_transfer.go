// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SetTransferHandlerFunc turns a function with the right signature into a set transfer handler
type SetTransferHandlerFunc func(SetTransferParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SetTransferHandlerFunc) Handle(params SetTransferParams) middleware.Responder {
	return fn(params)
}

// SetTransferHandler interface for that can handle valid set transfer params
type SetTransferHandler interface {
	Handle(SetTransferParams) middleware.Responder
}

// NewSetTransfer creates a new http.Handler for the set transfer operation
func NewSetTransfer(ctx *middleware.Context, handler SetTransferHandler) *SetTransfer {
	return &SetTransfer{Context: ctx, Handler: handler}
}

/*SetTransfer swagger:route POST /account/transfer Account setTransfer

Set a transfer

<b>Set a transfer</b>

*/
type SetTransfer struct {
	Context *middleware.Context
	Handler SetTransferHandler
}

func (o *SetTransfer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSetTransferParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
