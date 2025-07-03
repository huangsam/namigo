package core

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// DocumentGenerator generates a goquery Document.
type DocumentGenerator interface {
	// Execute generates the document.
	Execute() (*goquery.Document, error)
}

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
) DocumentGenerator {
	return &DocumentPipeline{
		httpClient:       client,
		requestBuilder:   builder,
		responseHandlers: handlers,
	}
}

// Execute executes the pipeline and returns the document.
func (dp *DocumentPipeline) Execute() (*goquery.Document, error) {
	req, err := dp.requestBuilder()
	if err != nil {
		return nil, err
	}
	res, err := dp.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer dismiss(res.Body.Close)
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
