package util

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type DocumentExecutor interface {
	// Execute generates a Document for goquery.
	Execute() (*goquery.Document, error)
}

// RequestBuilder builds a HTTP request for client.
type RequestBuilder func() (*http.Request, error)

// ResponseHandler handles side effects for a response.
type ResponseHandler func(*http.Response) error

// DocumentPipeline implements the DocumentExecutor interface.
type DocumentPipeline struct {
	httpClient       *http.Client      // HTTP client
	requestBuilder   RequestBuilder    // HTTP request builder
	responseHandlers []ResponseHandler // List of HTTP response handlers
}

// NewDocumentPipeline creates a pipeline for execution.
func NewDocumentPipeline(
	client *http.Client,
	builder RequestBuilder,
	handlers ...ResponseHandler,
) DocumentExecutor {
	return &DocumentPipeline{
		httpClient:       client,
		requestBuilder:   builder,
		responseHandlers: handlers,
	}
}

func (dp *DocumentPipeline) Execute() (*goquery.Document, error) {
	req, err := dp.requestBuilder()
	if err != nil {
		return nil, err
	}
	res, err := dp.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	for _, handler := range dp.responseHandlers { // For multiple side effects
		if err = handler(res); err != nil {
			return nil, err
		}
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
