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
	validate "github.com/go-openapi/validate"

	models "github.com/devchallenge/spy-api/internal/gen/models"
)

// PostBbinputHandlerFunc turns a function with the right signature into a post bbinput handler
type PostBbinputHandlerFunc func(PostBbinputParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostBbinputHandlerFunc) Handle(params PostBbinputParams) middleware.Responder {
	return fn(params)
}

// PostBbinputHandler interface for that can handle valid post bbinput params
type PostBbinputHandler interface {
	Handle(PostBbinputParams) middleware.Responder
}

// NewPostBbinput creates a new http.Handler for the post bbinput operation
func NewPostBbinput(ctx *middleware.Context, handler PostBbinputHandler) *PostBbinput {
	return &PostBbinput{Context: ctx, Handler: handler}
}

/*PostBbinput swagger:route POST /bbinput postBbinput

Accepts GPS coordinates from the mobile and saves them to the database

*/
type PostBbinput struct {
	Context *middleware.Context
	Handler PostBbinputHandler
}

func (o *PostBbinput) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostBbinputParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostBbinputBody post bbinput body
// swagger:model PostBbinputBody
type PostBbinputBody struct {

	// coordinates
	// Required: true
	Coordinates *PostBbinputParamsBodyCoordinates `json:"coordinates"`

	// Device identificator
	// Required: true
	Imei *string `json:"imei"`

	// Optional IP address
	IP string `json:"ip,omitempty"`

	// Phone
	// Required: true
	Number *string `json:"number"`

	// timestamp
	Timestamp models.Timestamp `json:"timestamp,omitempty"`
}

// Validate validates this post bbinput body
func (o *PostBbinputBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCoordinates(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImei(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostBbinputBody) validateCoordinates(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"coordinates", "body", o.Coordinates); err != nil {
		return err
	}

	if o.Coordinates != nil {
		if err := o.Coordinates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "coordinates")
			}
			return err
		}
	}

	return nil
}

func (o *PostBbinputBody) validateImei(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"imei", "body", o.Imei); err != nil {
		return err
	}

	return nil
}

func (o *PostBbinputBody) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"number", "body", o.Number); err != nil {
		return err
	}

	return nil
}

func (o *PostBbinputBody) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := o.Timestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "timestamp")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostBbinputBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBbinputBody) UnmarshalBinary(b []byte) error {
	var res PostBbinputBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostBbinputParamsBodyCoordinates GPS coordinates of the phone's location
// swagger:model PostBbinputParamsBodyCoordinates
type PostBbinputParamsBodyCoordinates struct {

	// Latitude in degrees
	// Maximum: 90
	// Minimum: -90
	Latitude *float32 `json:"latitude,omitempty"`

	// Longitude in degrees
	// Maximum: 180
	// Minimum: -180
	Longitude *float32 `json:"longitude,omitempty"`
}

// Validate validates this post bbinput params body coordinates
func (o *PostBbinputParamsBodyCoordinates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLatitude(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLongitude(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostBbinputParamsBodyCoordinates) validateLatitude(formats strfmt.Registry) error {

	if swag.IsZero(o.Latitude) { // not required
		return nil
	}

	if err := validate.Minimum("body"+"."+"coordinates"+"."+"latitude", "body", float64(*o.Latitude), -90, false); err != nil {
		return err
	}

	if err := validate.Maximum("body"+"."+"coordinates"+"."+"latitude", "body", float64(*o.Latitude), 90, false); err != nil {
		return err
	}

	return nil
}

func (o *PostBbinputParamsBodyCoordinates) validateLongitude(formats strfmt.Registry) error {

	if swag.IsZero(o.Longitude) { // not required
		return nil
	}

	if err := validate.Minimum("body"+"."+"coordinates"+"."+"longitude", "body", float64(*o.Longitude), -180, false); err != nil {
		return err
	}

	if err := validate.Maximum("body"+"."+"coordinates"+"."+"longitude", "body", float64(*o.Longitude), 180, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostBbinputParamsBodyCoordinates) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBbinputParamsBodyCoordinates) UnmarshalBinary(b []byte) error {
	var res PostBbinputParamsBodyCoordinates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
