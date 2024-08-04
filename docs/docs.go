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
        "/api/auth/forgot-password": {
            "post": {
                "description": "Request forgot password.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Forgot password.",
                "parameters": [
                    {
                        "description": "the body to request forgot password",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ForgotPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.WebSuccess-response_ForgotPasswordResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.WebNotFoundError"
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
        },
        "/api/auth/reset-password": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Reset password.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Reset password.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "the body to reset password",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ResetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.WebSuccess-string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.WebBadRequestError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.WebNotFoundError"
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
        "/api/transactions": {
            "get": {
                "description": "Registering a user from public access.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Get all transaction.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.WebSuccess-array_response_TransactionResponse"
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
        },
        "/api/transactions/payment/{id}": {
            "patch": {
                "description": "Pay for a transaction",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Pay for a transaction",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.WebSuccess-string"
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
        },
        "/api/transactions/{id}": {
            "get": {
                "description": "Get transaction by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Get transaction by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.WebSuccess-response_TransactionResponse"
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
            },
            "delete": {
                "description": "Delete a transaction",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Delete a transaction",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.WebSuccess-string"
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
            },
            "patch": {
                "description": "Update a transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Update a transaction",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Transaction update payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/request.TransactionUpdate"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.WebSuccess-string"
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
        },
        "/api/transactions/{userId}": {
            "post": {
                "description": "Create a new transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Create a new transaction",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Transaction payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/request.TransactionCreate"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.WebSuccess-string"
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
        },
        "/api/users/current": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Get current user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get current user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.WebSuccess-response_GetUserCurrentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.WebBadRequestError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.WebNotFoundError"
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
        "request.ForgotPasswordRequest": {
            "type": "object",
            "required": [
                "email",
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
                    "type": "string"
                }
            }
        },
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
        "request.ResetPasswordRequest": {
            "type": "object",
            "required": [
                "new_password"
            ],
            "properties": {
                "new_password": {
                    "type": "string",
                    "minLength": 8,
                    "x-order": "0",
                    "example": "new_password"
                }
            }
        },
        "request.TransactionCreate": {
            "type": "object",
            "properties": {
                "bikeId": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "totalPrice": {
                    "type": "integer"
                }
            }
        },
        "request.TransactionUpdate": {
            "type": "object",
            "properties": {
                "bikeId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "totalPrice": {
                    "type": "integer"
                }
            }
        },
        "response.ForgotPasswordResponse": {
            "type": "object",
            "properties": {
                "forgot_password_token": {
                    "type": "string",
                    "example": "token"
                }
            }
        },
        "response.GetUserCurrentResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "x-order": "0",
                    "example": 1
                },
                "username": {
                    "type": "string",
                    "x-order": "2",
                    "example": "luigi"
                },
                "email": {
                    "type": "string",
                    "x-order": "3",
                    "example": "luigi@sam.com"
                },
                "role": {
                    "type": "string",
                    "x-order": "4",
                    "example": "USER"
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
        "response.OrderResponse": {
            "type": "object",
            "properties": {
                "bikeID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "totalPrice": {
                    "type": "integer"
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
        "response.TransactionResponse": {
            "type": "object",
            "properties": {
                "autoCreateTime": {
                    "type": "string"
                },
                "autoUpdateTime": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.OrderResponse"
                    }
                },
                "status": {
                    "type": "string"
                },
                "totalPrice": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
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
        "web.WebNotFoundError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 404
                },
                "errors": {
                    "type": "string",
                    "example": "Not Found"
                }
            }
        },
        "web.WebSuccess-array_response_TransactionResponse": {
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
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.TransactionResponse"
                    },
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
        "web.WebSuccess-response_ForgotPasswordResponse": {
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
                            "$ref": "#/definitions/response.ForgotPasswordResponse"
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
        "web.WebSuccess-response_GetUserCurrentResponse": {
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
                            "$ref": "#/definitions/response.GetUserCurrentResponse"
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
        "web.WebSuccess-response_TransactionResponse": {
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
                            "$ref": "#/definitions/response.TransactionResponse"
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
        "web.WebSuccess-string": {
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
                    "type": "string",
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
