// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Spy API",
    "version": "1.0.0"
  },
  "basePath": "/ourell",
  "paths": {
    "/bbinput": {
      "post": {
        "description": "TODO",
        "produces": [
          "application/json"
        ],
        "summary": "TODO",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/bboutput": {
      "get": {
        "description": "TODO",
        "produces": [
          "application/json"
        ],
        "summary": "TODO",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Spy API",
    "version": "1.0.0"
  },
  "basePath": "/ourell",
  "paths": {
    "/bbinput": {
      "post": {
        "description": "TODO",
        "produces": [
          "application/json"
        ],
        "summary": "TODO",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/bboutput": {
      "get": {
        "description": "TODO",
        "produces": [
          "application/json"
        ],
        "summary": "TODO",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    }
  }
}`))
}
