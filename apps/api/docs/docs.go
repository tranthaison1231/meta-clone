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
        "/communities": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Get Communities",
                "operationId": "get-communities",
                "responses": {
                    "200": {
                        "description": "status\": \"success\", data: { \"communities\": []models.Community }, \"code\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/posts": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "summary": "Get Posts",
                "operationId": "get-posts",
                "responses": {
                    "200": {
                        "description": "status\": \"success\", data: { \"posts\": []models.Post }, \"code\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sign-in": {
            "post": {
                "summary": "Sign In",
                "operationId": "sign-in",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status\": \"success\", \"token\": string, \"code\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.SignInRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/dev",
	Schemes:          []string{},
	Title:            "Meta-Clone",
	Description:      "API for Meta-Clone",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
