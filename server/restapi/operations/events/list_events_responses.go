// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/5verst/server/models"
	"github.com/gadzira/5verst/internal/domain"
)

// ListEventsOKCode is the HTTP code returned for type ListEventsOK
const ListEventsOKCode int = 200

/*ListEventsOK Success response

swagger:response listEventsOK
*/
type ListEventsOK struct {

	/*
	  In: Body
	*/
	Payload domain.EventsList `json:"body,omitempty"`
}

// NewListEventsOK creates ListEventsOK with default headers values
func NewListEventsOK() *ListEventsOK {

	return &ListEventsOK{}
}

// WithPayload adds the payload to the list events o k response
func (o *ListEventsOK) WithPayload(payload domain.EventsList) *ListEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list events o k response
func (o *ListEventsOK) SetPayload(payload domain.EventsList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = domain.EventsList{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListEventsBadRequestCode is the HTTP code returned for type ListEventsBadRequest
const ListEventsBadRequestCode int = 400

/*ListEventsBadRequest (Bad Request) Invalid request parameters

swagger:response listEventsBadRequest
*/
type ListEventsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorBadRequest `json:"body,omitempty"`
}

// NewListEventsBadRequest creates ListEventsBadRequest with default headers values
func NewListEventsBadRequest() *ListEventsBadRequest {

	return &ListEventsBadRequest{}
}

// WithPayload adds the payload to the list events bad request response
func (o *ListEventsBadRequest) WithPayload(payload *models.ErrorBadRequest) *ListEventsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list events bad request response
func (o *ListEventsBadRequest) SetPayload(payload *models.ErrorBadRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEventsUnprocessableEntityCode is the HTTP code returned for type ListEventsUnprocessableEntity
const ListEventsUnprocessableEntityCode int = 422

/*ListEventsUnprocessableEntity Unprocessable entity

swagger:response listEventsUnprocessableEntity
*/
type ListEventsUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorUnprocessableEntity `json:"body,omitempty"`
}

// NewListEventsUnprocessableEntity creates ListEventsUnprocessableEntity with default headers values
func NewListEventsUnprocessableEntity() *ListEventsUnprocessableEntity {

	return &ListEventsUnprocessableEntity{}
}

// WithPayload adds the payload to the list events unprocessable entity response
func (o *ListEventsUnprocessableEntity) WithPayload(payload *models.ErrorUnprocessableEntity) *ListEventsUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list events unprocessable entity response
func (o *ListEventsUnprocessableEntity) SetPayload(payload *models.ErrorUnprocessableEntity) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventsUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
