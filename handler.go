package main

import (
	"fmt"
	"html"
	"log/slog"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slog.InfoContext(ctx, "home log...")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Api(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slog.InfoContext(ctx, "api log...")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func None(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slog.InfoContext(ctx, "none log...")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
