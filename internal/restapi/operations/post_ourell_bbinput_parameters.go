// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewPostOurellBbinputParams creates a new PostOurellBbinputParams object
// no default values defined in spec.
func NewPostOurellBbinputParams() PostOurellBbinputParams {

	return PostOurellBbinputParams{}
}

// PostOurellBbinputParams contains all the bound params for the post ourell bbinput operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostOurellBbinput
type PostOurellBbinputParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostOurellBbinputParams() beforehand.
func (o *PostOurellBbinputParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
