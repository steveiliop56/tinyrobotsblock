package tinyrobotsblock

import (
	"context"
	"net/http"
	"strings"
)

// Config

type Config struct {
	Enabled bool `json:"enabled"`
}

func CreateConfig() *Config {
	return &Config{
		Enabled: true,
	}
}

// Plugin

type Tinyrobotsblock struct {
	next    http.Handler
	name    string
	enabled bool
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Tinyrobotsblock{
		next:    next,
		name:    name,
		enabled: config.Enabled,
	}, nil
}

func (t *Tinyrobotsblock) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix(req.URL.Path, "/robots.txt") && t.enabled {
		rw.Header().Set("Content-Type", "text/plain")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("User-agent: *\nDisallow: /\n"))
		return
	}
	t.next.ServeHTTP(rw, req)
}
