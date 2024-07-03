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
    "description": "API for retrieving user data",
    "title": "User API",
    "version": "1.0.0"
  },
  "host": "localhost:9091",
  "basePath": "/v1",
  "paths": {
    "/users": {
      "get": {
        "description": "Retrieve a list of users with optional filtering.",
        "produces": [
          "application/json"
        ],
        "summary": "Get Users",
        "parameters": [
          {
            "type": "string",
            "description": "Filter users by email",
            "name": "email",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A list of users",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/User"
              }
            }
          },
          "400": {
            "description": "Invalid parameter"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "first_name": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "last_name": {
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
    "description": "API for retrieving user data",
    "title": "User API",
    "version": "1.0.0"
  },
  "host": "localhost:9091",
  "basePath": "/v1",
  "paths": {
    "/users": {
      "get": {
        "description": "Retrieve a list of users with optional filtering.",
        "produces": [
          "application/json"
        ],
        "summary": "Get Users",
        "parameters": [
          {
            "type": "string",
            "description": "Filter users by email",
            "name": "email",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A list of users",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/User"
              }
            }
          },
          "400": {
            "description": "Invalid parameter"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "first_name": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "last_name": {
          "type": "string"
        }
      }
    }
  }
}`))
}