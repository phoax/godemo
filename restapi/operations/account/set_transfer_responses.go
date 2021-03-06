// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/phoax/godemo/models"
)

// SetTransferCreatedCode is the HTTP code returned for type SetTransferCreated
const SetTransferCreatedCode int = 201

/*SetTransferCreated transfer

swagger:response setTransferCreated
*/
type SetTransferCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Ack `json:"body,omitempty"`
}

// NewSetTransferCreated creates SetTransferCreated with default headers values
func NewSetTransferCreated() *SetTransferCreated {

	return &SetTransferCreated{}
}

// WithPayload adds the payload to the set transfer created response
func (o *SetTransferCreated) WithPayload(payload *models.Ack) *SetTransferCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set transfer created response
func (o *SetTransferCreated) SetPayload(payload *models.Ack) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetTransferCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*SetTransferDefault Internal error

swagger:response setTransferDefault
*/
type SetTransferDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetTransferDefault creates SetTransferDefault with default headers values
func NewSetTransferDefault(code int) *SetTransferDefault {
	if code <= 0 {
		code = 500
	}

	return &SetTransferDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the set transfer default response
func (o *SetTransferDefault) WithStatusCode(code int) *SetTransferDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the set transfer default response
func (o *SetTransferDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the set transfer default response
func (o *SetTransferDefault) WithPayload(payload *models.Error) *SetTransferDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set transfer default response
func (o *SetTransferDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetTransferDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
