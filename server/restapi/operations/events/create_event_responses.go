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

// CreateEventCreatedCode is the HTTP code returned for type CreateEventCreated
const CreateEventCreatedCode int = 201

/*CreateEventCreated Success response

swagger:response createEventCreated
*/
type CreateEventCreated struct {

	/*
	  In: Body
	*/
	Payload *domain.Event `json:"body,omitempty"`
}

// NewCreateEventCreated creates CreateEventCreated with default headers values
func NewCreateEventCreated() *CreateEventCreated {

	return &CreateEventCreated{}
}

// WithPayload adds the payload to the create event created response
func (o *CreateEventCreated) WithPayload(payload *domain.Event) *CreateEventCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create event created response
func (o *CreateEventCreated) SetPayload(payload *domain.Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEventCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEventBadRequestCode is the HTTP code returned for type CreateEventBadRequest
const CreateEventBadRequestCode int = 400

/*CreateEventBadRequest (Bad Request) Invalid request parameters

swagger:response createEventBadRequest
*/
type CreateEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorBadRequest `json:"body,omitempty"`
}

// NewCreateEventBadRequest creates CreateEventBadRequest with default headers values
func NewCreateEventBadRequest() *CreateEventBadRequest {

	return &CreateEventBadRequest{}
}

// WithPayload adds the payload to the create event bad request response
func (o *CreateEventBadRequest) WithPayload(payload *models.ErrorBadRequest) *CreateEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create event bad request response
func (o *CreateEventBadRequest) SetPayload(payload *models.ErrorBadRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEventUnprocessableEntityCode is the HTTP code returned for type CreateEventUnprocessableEntity
const CreateEventUnprocessableEntityCode int = 422

/*CreateEventUnprocessableEntity Unprocessable entity

swagger:response createEventUnprocessableEntity
*/
type CreateEventUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorUnprocessableEntity `json:"body,omitempty"`
}

// NewCreateEventUnprocessableEntity creates CreateEventUnprocessableEntity with default headers values
func NewCreateEventUnprocessableEntity() *CreateEventUnprocessableEntity {

	return &CreateEventUnprocessableEntity{}
}

// WithPayload adds the payload to the create event unprocessable entity response
func (o *CreateEventUnprocessableEntity) WithPayload(payload *models.ErrorUnprocessableEntity) *CreateEventUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create event unprocessable entity response
func (o *CreateEventUnprocessableEntity) SetPayload(payload *models.ErrorUnprocessableEntity) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEventUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
