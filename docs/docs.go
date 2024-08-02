// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/login": {
            "post": {
                "description": "Logging in to get jwt token to access admin or user api by roles.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User login.",
                "parameters": [
                    {
                        "description": "the body to login a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.WebSuccess-response_LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.WebBadRequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/web.WebUnauthorizedError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.WebInternalServerError"
                        }
                    }
                }
            }
        },
        "/api/auth/register": {
            "post": {
                "description": "Registering a user from public access.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User register.",
                "parameters": [
                    {
                        "description": "the body to register a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/web.WebSuccess-response_RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.WebBadRequestError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.WebInternalServerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "x-order": "0"
                },
                "password": {
                    "type": "string",
                    "x-order": "1",
                    "example": "password"
                }
            }
        },
        "request.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "username": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3,
                    "x-order": "0"
                },
                "email": {
                    "type": "string",
                    "x-order": "1"
                },
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "x-order": "2",
                    "example": "password"
                }
            }
        },
        "response.LoginResponse": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string",
                    "x-order": "0",
                    "example": "luigi"
                },
                "email": {
                    "type": "string",
                    "x-order": "1",
                    "example": "luigi@sam.com"
                },
                "role": {
                    "type": "string",
                    "x-order": "2",
                    "example": "USER"
                },
                "token": {
                    "type": "string",
                    "x-order": "3",
                    "example": "token"
                }
            }
        },
        "response.RegisterResponse": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string",
                    "x-order": "0",
                    "example": "luigi"
                },
                "email": {
                    "type": "string",
                    "x-order": "1",
                    "example": "luigi@sam.com"
                },
                "role": {
                    "type": "string",
                    "x-order": "2",
                    "example": "USER"
                }
            }
        },
        "web.Metadata": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer",
                    "x-order": "0"
                },
                "limit": {
                    "type": "integer",
                    "x-order": "1"
                },
                "total_pages": {
                    "type": "integer",
                    "x-order": "2"
                },
                "total_data": {
                    "type": "integer",
                    "x-order": "3"
                }
            }
        },
        "web.WebBadRequestError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "errors": {
                    "type": "string",
                    "example": "Bad Request"
                }
            }
        },
        "web.WebInternalServerError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "errors": {
                    "type": "string",
                    "example": "Internal Server Error"
                }
            }
        },
        "web.WebSuccess-response_LoginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "x-order": "0",
                    "example": 200
                },
                "message": {
                    "type": "string",
                    "x-order": "1",
                    "example": "success"
                },
                "payload": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/response.LoginResponse"
                        }
                    ],
                    "x-order": "2"
                },
                "metadata": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/web.Metadata"
                        }
                    ],
                    "x-order": "3"
                }
            }
        },
        "web.WebSuccess-response_RegisterResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "x-order": "0",
                    "example": 200
                },
                "message": {
                    "type": "string",
                    "x-order": "1",
                    "example": "success"
                },
                "payload": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/response.RegisterResponse"
                        }
                    ],
                    "x-order": "2"
                },
                "metadata": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/web.Metadata"
                        }
                    ],
                    "x-order": "3"
                }
            }
        },
        "web.WebUnauthorizedError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 401
                },
                "errors": {
                    "type": "string",
                    "example": "Unauthorized"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
