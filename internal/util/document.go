package util

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type DocumentExecutor interface {
	// Execute generates a Document for goquery.
	Execute() (*goquery.Document, error)
}

type RequestBuilder func() (*http.Request, error)

type ResponseHandler func(*http.Response) error

// DocumentPipeline implements the DocumentExecutor interface.
type DocumentPipeline struct {
	client   *http.Client      // HTTP client
	builder  RequestBuilder    // HTTP request builder
	handlers []ResponseHandler // List of HTTP response handlers
}

// NewDocumentPipeline creates a pipeline for execution.
func NewDocumentPipeline(client *http.Client, builder RequestBuilder, handlers ...ResponseHandler) DocumentExecutor {
	return &DocumentPipeline{
		client:   client,
		builder:  builder,
		handlers: handlers,
	}
}

func (dp *DocumentPipeline) Execute() (*goquery.Document, error) {
	req, err := dp.builder()
	if err != nil {
		return nil, err
	}
	res, err := dp.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	for _, handler := range dp.handlers { // For multiple side effects
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
