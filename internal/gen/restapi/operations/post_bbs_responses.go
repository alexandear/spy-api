// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/devchallenge/spy-api/internal/gen/models"
)

// PostBbsOKCode is the HTTP code returned for type PostBbsOK
const PostBbsOKCode int = 200

/*PostBbsOK OK

swagger:response postBbsOK
*/
type PostBbsOK struct {

	/*
	  In: Body
	*/
	Payload *PostBbsOKBody `json:"body,omitempty"`
}

// NewPostBbsOK creates PostBbsOK with default headers values
func NewPostBbsOK() *PostBbsOK {

	return &PostBbsOK{}
}

// WithPayload adds the payload to the post bbs o k response
func (o *PostBbsOK) WithPayload(payload *PostBbsOKBody) *PostBbsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post bbs o k response
func (o *PostBbsOK) SetPayload(payload *PostBbsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBbsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostBbsBadRequestCode is the HTTP code returned for type PostBbsBadRequest
const PostBbsBadRequestCode int = 400

/*PostBbsBadRequest Invalid arguments

swagger:response postBbsBadRequest
*/
type PostBbsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostBbsBadRequest creates PostBbsBadRequest with default headers values
func NewPostBbsBadRequest() *PostBbsBadRequest {

	return &PostBbsBadRequest{}
}

// WithPayload adds the payload to the post bbs bad request response
func (o *PostBbsBadRequest) WithPayload(payload *models.Error) *PostBbsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post bbs bad request response
func (o *PostBbsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBbsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostBbsInternalServerErrorCode is the HTTP code returned for type PostBbsInternalServerError
const PostBbsInternalServerErrorCode int = 500

/*PostBbsInternalServerError General server error

swagger:response postBbsInternalServerError
*/
type PostBbsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostBbsInternalServerError creates PostBbsInternalServerError with default headers values
func NewPostBbsInternalServerError() *PostBbsInternalServerError {

	return &PostBbsInternalServerError{}
}

// WithPayload adds the payload to the post bbs internal server error response
func (o *PostBbsInternalServerError) WithPayload(payload *models.Error) *PostBbsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post bbs internal server error response
func (o *PostBbsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBbsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
