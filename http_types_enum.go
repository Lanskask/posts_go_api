package main

type HttpResponseType string

const (
	ApplicationJson   HttpResponseType = "application/json"
	ApplicationXml                     = "application/xml"
	MultipartFormData                  = "multipart/form-data"
	TextHtml                           = "text/html"
	TextPlain                          = "text/plain"
	TextXml                            = "text/xml"
	Wildcard                           = "*/*"
)

// "Content-Type"."application/json"
