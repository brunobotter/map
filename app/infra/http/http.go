package http

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	_http "net/http"
	"net/url"
	"time"

	"github.com/brunobotter/map/application/http"
	"github.com/brunobotter/map/infra/logger"
	"github.com/brunobotter/map/main/config"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

const (
	requestTimeout = 5 * time.Second
	maxRetries     = 2
	retryDelay     = 200 * time.Millisecond
)

type httpClient struct {
	cfg    *config.Config
	client *_http.Client
	logger logger.Logger
}

func NewHttpClient(cfg *config.Config, logger logger.Logger) http.Client {
	dialer := new(net.Dialer)

	transport := new(_http.Transport)
	transport.DialContext = dialer.DialContext
	transport.MaxIdleConnsPerHost = 5

	if cfg.Env == "local" {
		transport.Proxy = func(req *_http.Request) (*url.URL, error) {
			return _http.ProxyFromEnvironment(req)
		}
	}

	client := new(_http.Client)
	client.Transport = transport
	client.Timeout = requestTimeout
	client = WrapClient(client)
	return &httpClient{cfg, client, logger}
}

func WrapClient(client *_http.Client) *_http.Client {
	return httptrace.WrapClient(client,
		httptrace.RTWithResourceNamer(func(req *_http.Request) string {
			return fmt.Sprintf("%s - %s", req.Method, req.Host)
		}),
	)
}

func (c *httpClient) NewRequestWithContext(ctx context.Context, method string, url string, body []byte) (request http.Request, err error) {
	request, err = newRequestWithContext(ctx, method, url, body)
	if err != nil {
		return request, err
	}
	return request, nil
}

func (c *httpClient) NewRequest(method string, url string, body []byte) (request http.Request, err error) {
	request, err = newRequest(method, url, body)
	if err != nil {
		return request, err
	}
	return request, nil
}

func (c *httpClient) Do(ctx context.Context, service string, req http.Request) (response http.Response, err error) {
	request, ok := req.(*request)
	if !ok {
		return response, errors.New("invalid request")
	}
	_ = service

	for attempt := 0; attempt <= maxRetries; attempt++ {
		httpReq := request.Req.Clone(ctx)
		if request.Body != nil {
			httpReq.Body = io.NopCloser(bytes.NewReader(request.Body))
		}

		attemptCtx, cancel := context.WithTimeout(httpReq.Context(), requestTimeout)
		httpReq = httpReq.WithContext(attemptCtx)

		resp, doErr := c.client.Do(httpReq)
		cancel()
		if doErr != nil {
			if attempt == maxRetries || !shouldRetryError(doErr) {
				return response, doErr
			}

			if err = waitRetry(ctx); err != nil {
				return response, err
			}
			continue
		}

		if resp.StatusCode >= _http.StatusInternalServerError && attempt < maxRetries {
			resp.Body.Close()
			if err = waitRetry(ctx); err != nil {
				return response, err
			}
			continue
		}

		defer resp.Body.Close()
		body, bodyReadErr := io.ReadAll(resp.Body)
		return newResponse(resp, body, bodyReadErr), nil
	}

	return response, errors.New("request failed after retries")
}

func shouldRetryError(err error) bool {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		if netErr.Timeout() {
			return false
		}
		return netErr.Temporary()
	}

	return false
}

func waitRetry(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(retryDelay):
		return nil
	}
}
