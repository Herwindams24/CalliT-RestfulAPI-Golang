package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": ["http"],
    "swagger": "2.0",
    "info": {
        "title": "Panggilan Darurat Rumah Sakit API",
    },
    "host": "callit.lintasin.me",
    "basePath": "/",
    "securityDefinitions": {        
        "bearerAuth": {
            "type": "apiKey",
            "scheme": "bearer",
            "name": "Authorization",
            "in": "header"
        }
      },
    "paths": {
        "/daftarInstansi": {
            "get": {
                "description": "Data Penulis yang terdaftar pada iLib",
                "tags" : [daftarInstansi],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Daftar Panggilan darurat Rumah Sakit",
                "parameters": [],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string",
                            "$ref": "#/definitions/web.GetdaftarInstansiSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema" : {
                            "type": "string",
                            "$ref": "#/definitions/web.WrongMethod"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "post": {
                "description": "Pada metode ini pengguna dapat menambahkan Daftar Panggilan darurat Rumah Sakit",
                "tags" : [daftarInstansi],
                "consumes": [
                    "application/json"
                ],
                "security":[
                    "bearerAuth": []
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Menambahkan Daftar Panggilan darurat Rumah Sakit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namaInstansi",
                        "name": "namaInstansi",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Kota",
                        "name": "Kota",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "telfon",
                        "name": "telfon",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string",
                            "$ref": "#/definitions/web.PostdaftarInstansiSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/web.WrongMethod"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "description": "Di metode ini user dapat mengedit nama, kota, maupun telepon Rumah Sakit.",
                "tags" : [daftarInstansi],
                "consumes": [
                    "application/json"
                ],
                "security":[
                    "bearerAuth": []
                ],
                "produces": [
                    "application/json"
                ],  
                "summary": "Menyunting data yang telah ada",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namaInstansi",
                        "name": "namaInstansi",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Kota",
                        "name": "Kota",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "telfon",
                        "name": "telfon",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string",
                            "$ref": "#/definitions/web.PutdaftarInstansiSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/web.WrongMethod"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "delete": {
                "description": "Menghapus data instansi yang ada di API",
                "tags" : [daftarInstansi],
                "consumes": [
                    "application/json"
                ],
                "security":[
                    "bearerAuth": []
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Melakukan penghapusan data Rumah Sakit yang tidak relevan",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string",
                            "$ref": "#/definitions/web.DeletedaftarInstansiSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/web.WrongMethod"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        }
    },
    "definitions": {
        "web.GetdaftarInstansiSuccess": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "nama": {
                    "type": "string"
                }
            }
        },
        "web.PostdaftarInstansiSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                }
            }
        },
        "web.PutdaftarInstansiSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                }
            }
        },
        "web.DeletedaftarInstansiSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                }
            }
        },
        "web.WrongMethod": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
