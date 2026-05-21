package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestIsLandingPageRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want bool
	}{
		{
			name: "origin form get",
			req: &http.Request{
				Method: http.MethodGet,
				URL:    &url.URL{Path: "/"},
				Host:   "web.nacl.one",
			},
			want: true,
		},
		{
			name: "origin form head",
			req: &http.Request{
				Method: http.MethodHead,
				URL:    &url.URL{Path: "/"},
				Host:   "web.nacl.one",
			},
			want: true,
		},
		{
			name: "absolute form get",
			req: &http.Request{
				Method: http.MethodGet,
				URL:    &url.URL{Scheme: "http", Host: "openrouter.ai", Path: "/"},
				Host:   "openrouter.ai",
			},
			want: false,
		},
		{
			name: "connect",
			req: &http.Request{
				Method: http.MethodConnect,
				URL:    &url.URL{Host: "openrouter.ai:443"},
				Host:   "openrouter.ai:443",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLandingPageRequest(tt.req); got != tt.want {
				t.Fatalf("isLandingPageRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServeLandingPage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "https://web.nacl.one/", nil)
	rr := httptest.NewRecorder()

	serveLandingPage(rr, req)

	resp := rr.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if got := resp.Header.Get("Content-Type"); got != "text/html; charset=utf-8" {
		t.Fatalf("content-type = %q", got)
	}
	body := rr.Body.String()
	if !strings.Contains(body, "NaCl") || !strings.Contains(body, "sodium chloride") {
		t.Fatalf("body does not contain expected landing page text: %q", body)
	}
}
