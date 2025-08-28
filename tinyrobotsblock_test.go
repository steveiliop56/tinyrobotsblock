package tinyrobotsblock_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/steveiliop56/tinyrobotsblock"
)

func TestTinyrobotsblock(t *testing.T) {
	cfg := tinyrobotsblock.CreateConfig()
	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := tinyrobotsblock.New(ctx, next, cfg, "tinyrobotsblock")

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/robots.txt", nil)

	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	expectedBody := "User-agent: *\nDisallow: /\n"

	if recorder.Body.String() != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, recorder.Body.String())
	}

	recorder = httptest.NewRecorder()

	req, err = http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/someotherpath", nil)

	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	if recorder.Body.String() != "" {
		t.Errorf("Expected empty body, got %q", recorder.Body.String())
	}
}
