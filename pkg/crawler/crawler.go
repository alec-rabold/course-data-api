package crawler

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Crawler abstracts
type Crawler struct {
	client  http.Client
	headers http.Header
}

type option func(*Crawler)

// NewCrawler returns a default crawler with the given options
func NewCrawler(opts ...option) *Crawler {
	c := defaultCrawler()
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func defaultCrawler() *Crawler {
	return &Crawler{
		headers: make(http.Header),
		client: http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:        100,
				IdleConnTimeout:     90 * time.Second,
				TLSHandshakeTimeout: 10 * time.Second,

				ExpectContinueTimeout: 4 * time.Second,
				ResponseHeaderTimeout: 5 * time.Second,
			},
		},
	}
}

func (c *Crawler) applyDefaults(request *http.Request) {
	for key, values := range c.headers {
		for _, value := range values {
			request.Header.Set(key, value)
		}
	}
}

// WithContentType applies a default content-type header
func WithContentType(contentType string) option {
	return func(c *Crawler) {
		c.headers.Set("Content-Type", contentType)
	}
}

// WithHeaders applies a set of global headers
func WithHeaders(headers map[string]string) option {
	return func(c *Crawler) {
		for key, value := range headers {
			c.headers.Set(key, value)
		}
	}
}

// WithTransport sets the provided transport
func WithTransport(transport *http.Transport) option {
	return func(c *Crawler) {
		c.client.Transport = transport
	}
}

// GetDocumentOption applies a configuration to the GetDocument request
type GetDocumentOption func(*http.Request)

// GetDocument translates the requested content into a parseable document
func (c *Crawler) GetDocument(ctx context.Context, url string, opts ...GetDocumentOption) (*goquery.Document, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("crafting request: %w", err)
	}
	// Apply common config
	c.applyDefaults(request)
	// Apply per-request options
	for _, opt := range opts {
		opt(request)
	}

	resp, err := c.client.Do(request)
	if err != nil {
		if resp != nil {
			return nil, fmt.Errorf("getting document (HTTP %d %s): %v", resp.StatusCode, resp.Status, err)
		}
		return nil, fmt.Errorf("getting document: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing document from body: %w", err)
	}

	return doc, nil
}

// WithData applies the given data to the request body
func WithData(data url.Values) func(*http.Request) {
	return func(request *http.Request) {
		request.Body = io.NopCloser(strings.NewReader(data.Encode()))
	}
}

// WithReferer adds a referer header to the request headers
func WithReferer(referer string) func(*http.Request) {
	return func(request *http.Request) {
		request.Header.Set("Referer", referer)
	}
}
