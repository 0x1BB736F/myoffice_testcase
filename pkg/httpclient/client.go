package httpclient

import (
	"context"
	"io"

	"github.com/pkg/errors"

	"net/http"
	"time"

	urlverifier "github.com/davidmytton/url-verifier"
)

var (
	ErrNotOK = errors.New("http response is not OK")

	ErrNotValidScheme = errors.New("not valid scheme")
)

type HttpClient struct {
	cl       *http.Client
	verifier *urlverifier.Verifier
}

func New(timeout time.Duration) *HttpClient {
	return &HttpClient{
		cl: &http.Client{
			Timeout: timeout,
		},
		verifier: urlverifier.NewVerifier(),
	}
}

func (c HttpClient) Get(ctx context.Context, URL string) (HttpResponse, error) {
	handleStart := time.Now()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	if err != nil {
		return HttpResponse{}, err
	}

	res, err := c.cl.Do(req)
	if err != nil {
		return HttpResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return HttpResponse{}, ErrNotOK
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return HttpResponse{}, err
	}

	return HttpResponse{
		body:       b,
		handleTime: time.Since(handleStart),
	}, nil
}

func (c HttpClient) VerifyURL(URL string) error {
	ret, err := c.verifier.Verify(URL)
	if err != nil {
		return err
	}

	if ret.URLComponents != nil {
		scheme := ret.URLComponents.Scheme
		if scheme != "https" &&
			scheme != "http" {
			return ErrNotValidScheme
		}
	}

	return nil
}

type HttpResponse struct {
	body       []byte
	handleTime time.Duration
}

func (r HttpResponse) ContentLength() int {
	return len(r.body)
}

func (r HttpResponse) HandleTime() time.Duration {
	return r.handleTime
}
