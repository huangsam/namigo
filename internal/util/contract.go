package util

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// RequestBuilder builds a HTTP request for client.
type RequestBuilder func() (*http.Request, error)

// ResponseHandler handles side effects for a response.
type ResponseHandler func(*http.Response) error

// DocumentGenerator generates a goquery Document.
type DocumentGenerator interface {
	// Execute generates the document.
	Execute() (*goquery.Document, error)
}
