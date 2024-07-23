// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/metrics": {
            "get": {
                "description": "Consult quote metrics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "metrics"
                ],
                "summary": "Metrics",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Number of quotes (descending order)",
                        "name": "last_quotes",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Metrics"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ControllerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ControllerError"
                        }
                    }
                }
            }
        },
        "/v1/quote": {
            "post": {
                "description": "Route for receiving input data and generating a freight quote",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quote"
                ],
                "summary": "Quote",
                "parameters": [
                    {
                        "description": "quote request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.QuoteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Quote"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ControllerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ControllerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_henrique77_api-quote_model_controller.Recipient": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/model.Address"
                }
            }
        },
        "github_com_henrique77_api-quote_model_controller.Volume": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "category": {
                    "type": "integer"
                },
                "height": {
                    "type": "number"
                },
                "length": {
                    "type": "number"
                },
                "price": {
                    "type": "number"
                },
                "sku": {
                    "type": "string"
                },
                "unitary_weight": {
                    "type": "number"
                },
                "width": {
                    "type": "number"
                }
            }
        },
        "model.Address": {
            "type": "object",
            "required": [
                "zipcode"
            ],
            "properties": {
                "zipcode": {
                    "type": "string"
                }
            }
        },
        "model.ControllerError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "model.Metrics": {
            "type": "object",
            "properties": {
                "average_final_price": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "number"
                    }
                },
                "least_expensive_shipping": {
                    "type": "number"
                },
                "most_expensive_shipping": {
                    "type": "number"
                },
                "results_per_carrier": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "total_final_price": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "number"
                    }
                }
            }
        },
        "model.Quote": {
            "type": "object",
            "properties": {
                "deadline": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "service": {
                    "type": "string"
                }
            }
        },
        "model.QuoteRequest": {
            "type": "object",
            "properties": {
                "recipient": {
                    "$ref": "#/definitions/github_com_henrique77_api-quote_model_controller.Recipient"
                },
                "volumes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_henrique77_api-quote_model_controller.Volume"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "API Quote",
	Description:      "API responsible for managing freight quotes",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
