// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetOurellBboutputOKCode is the HTTP code returned for type GetOurellBboutputOK
const GetOurellBboutputOKCode int = 200

/*GetOurellBboutputOK OK

swagger:response getOurellBboutputOK
*/
type GetOurellBboutputOK struct {
}

// NewGetOurellBboutputOK creates GetOurellBboutputOK with default headers values
func NewGetOurellBboutputOK() *GetOurellBboutputOK {

	return &GetOurellBboutputOK{}
}

// WriteResponse to the client
func (o *GetOurellBboutputOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
