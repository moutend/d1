package transport

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Transport struct {
	debug *log.Logger
}

func New() *Transport {
	return &Transport{
		debug: log.New(io.Discard, "", 0),
	}
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	now := time.Now().UTC()

	t.debug.Printf("request: sent at: %s\n", now.Format(time.RFC3339))
	t.debug.Printf("request: url: %s %s\n", req.Method, req.URL)

	req.Header.Set("User-Agent", "https://github.com/moutend/d1")

	for k, v := range req.Header {
		t.debug.Printf("request: header: %s: %s", k, strings.Join(v, ";"))
	}

	res, err := http.DefaultTransport.RoundTrip(req)

	if err != nil {
		return nil, fmt.Errorf("transport: %w", err)
	}

	defer res.Body.Close()

	buffer := &bytes.Buffer{}

	body, err := io.ReadAll(io.TeeReader(res.Body, buffer))

	if err != nil {
		return nil, fmt.Errorf("transport: %w", err)
	}

	res.Body = io.NopCloser(buffer)

	t.debug.Printf("response: status: %s\n", res.Status)

	for k, v := range res.Header {
		t.debug.Printf("response: header: %s: %s", k, strings.Join(v, ";"))
	}

	t.debug.Printf("response: body: %q\n", body)

	return res, nil
}

func (t *Transport) SetLogger(l *log.Logger) {
	if l == nil {
		return
	}

	t.debug = l
}
