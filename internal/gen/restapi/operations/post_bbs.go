// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/devchallenge/spy-api/internal/gen/models"
)

// PostBbsHandlerFunc turns a function with the right signature into a post bbs handler
type PostBbsHandlerFunc func(PostBbsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostBbsHandlerFunc) Handle(params PostBbsParams) middleware.Responder {
	return fn(params)
}

// PostBbsHandler interface for that can handle valid post bbs params
type PostBbsHandler interface {
	Handle(PostBbsParams) middleware.Responder
}

// NewPostBbs creates a new http.Handler for the post bbs operation
func NewPostBbs(ctx *middleware.Context, handler PostBbsHandler) *PostBbs {
	return &PostBbs{Context: ctx, Handler: handler}
}

/*PostBbs swagger:route POST /bbs postBbs

Shows how much time two phones are located in the same room

*/
type PostBbs struct {
	Context *middleware.Context
	Handler PostBbsHandler
}

func (o *PostBbs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostBbsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostBbsBody post bbs body
// swagger:model PostBbsBody
type PostBbsBody struct {

	// from
	// Required: true
	From models.Timestamp `json:"from"`

	// Distance in meters
	MinDistance int32 `json:"minDistance,omitempty"`

	// number1
	// Required: true
	Number1 models.Number `json:"number1"`

	// number2
	// Required: true
	Number2 models.Number `json:"number2"`

	// to
	// Required: true
	To models.Timestamp `json:"to"`
}

// Validate validates this post bbs body
func (o *PostBbsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNumber1(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNumber2(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostBbsBody) validateFrom(formats strfmt.Registry) error {

	if err := o.From.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "from")
		}
		return err
	}

	return nil
}

func (o *PostBbsBody) validateNumber1(formats strfmt.Registry) error {

	if err := o.Number1.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "number1")
		}
		return err
	}

	return nil
}

func (o *PostBbsBody) validateNumber2(formats strfmt.Registry) error {

	if err := o.Number2.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "number2")
		}
		return err
	}

	return nil
}

func (o *PostBbsBody) validateTo(formats strfmt.Registry) error {

	if err := o.To.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "to")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostBbsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBbsBody) UnmarshalBinary(b []byte) error {
	var res PostBbsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostBbsOKBody post bbs o k body
// swagger:model PostBbsOKBody
type PostBbsOKBody struct {

	// The number of off-hours (excluding the interval from 9 to 18) that people spend together
	Percentage int32 `json:"percentage,omitempty"`
}

// Validate validates this post bbs o k body
func (o *PostBbsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostBbsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBbsOKBody) UnmarshalBinary(b []byte) error {
	var res PostBbsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}