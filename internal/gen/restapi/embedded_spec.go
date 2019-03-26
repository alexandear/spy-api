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
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Monitoring for the mobile phones' moves",
    "title": "Spy API",
    "version": "1.0.0"
  },
  "basePath": "/ourell",
  "paths": {
    "/bbinput": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "summary": "Accepts GPS coordinates from the mobile and saves them to the database",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "number",
                "imei",
                "coordinates"
              ],
              "properties": {
                "coordinates": {
                  "description": "GPS coordinates of the phone's location",
                  "type": "object",
                  "properties": {
                    "latitude": {
                      "description": "Latitude in degrees",
                      "type": "number",
                      "format": "float",
                      "maximum": 90,
                      "minimum": -90
                    },
                    "longitude": {
                      "description": "Longitude in degrees",
                      "type": "number",
                      "format": "float",
                      "maximum": 180,
                      "minimum": -180
                    }
                  }
                },
                "imei": {
                  "description": "Device identificator",
                  "type": "string"
                },
                "ip": {
                  "description": "Optional IP address",
                  "type": "string"
                },
                "number": {
                  "description": "Phone",
                  "type": "string"
                },
                "timestamp": {
                  "description": "EET timestamp in \"YYYY/MM/DD-hh:mm:ss\" format",
                  "type": "string"
                }
              },
              "example": {
                "coordinates": [
                  22.1832284135991,
                  60.4538416572538
                ],
                "imei": "502507345219189",
                "ip": "35.25.21.123",
                "number": "+380991926482",
                "timestamp": "2019/03/22-15:50:20"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid arguments",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "General server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Monitoring for the mobile phones' moves",
    "title": "Spy API",
    "version": "1.0.0"
  },
  "basePath": "/ourell",
  "paths": {
    "/bbinput": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "summary": "Accepts GPS coordinates from the mobile and saves them to the database",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "number",
                "imei",
                "coordinates"
              ],
              "properties": {
                "coordinates": {
                  "description": "GPS coordinates of the phone's location",
                  "type": "object",
                  "properties": {
                    "latitude": {
                      "description": "Latitude in degrees",
                      "type": "number",
                      "format": "float",
                      "maximum": 90,
                      "minimum": -90
                    },
                    "longitude": {
                      "description": "Longitude in degrees",
                      "type": "number",
                      "format": "float",
                      "maximum": 180,
                      "minimum": -180
                    }
                  }
                },
                "imei": {
                  "description": "Device identificator",
                  "type": "string"
                },
                "ip": {
                  "description": "Optional IP address",
                  "type": "string"
                },
                "number": {
                  "description": "Phone",
                  "type": "string"
                },
                "timestamp": {
                  "description": "EET timestamp in \"YYYY/MM/DD-hh:mm:ss\" format",
                  "type": "string"
                }
              },
              "example": {
                "coordinates": [
                  22.1832284135991,
                  60.4538416572538
                ],
                "imei": "502507345219189",
                "ip": "35.25.21.123",
                "number": "+380991926482",
                "timestamp": "2019/03/22-15:50:20"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid arguments",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "General server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
}