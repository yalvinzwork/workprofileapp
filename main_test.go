package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "root path returns 200",
			path:           "/",
			expectedStatus: http.StatusOK,
			expectedBody:   "Welcome to my work profile!",
		},
		{
			name:           "invalid path returns 404",
			path:           "/invalid",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 page not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			w := httptest.NewRecorder()

			helloHandler(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("got status %d, want %d", w.Code, tt.expectedStatus)
			}

			if !strings.Contains(w.Body.String(), tt.expectedBody) {
				t.Errorf("got body %q, want to contain %q", w.Body.String(), tt.expectedBody)
			}
		})
	}
}

func TestAboutMeHandler(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "about-me path returns 200",
			path:           "/about-me",
			expectedStatus: http.StatusOK,
			expectedBody:   "I am an Infrastructure/SRE/DevOps professional",
		},
		{
			name:           "invalid about path returns 404",
			path:           "/about",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 page not found",
		},
		{
			name:           "root path in about handler returns 404",
			path:           "/",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 page not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			w := httptest.NewRecorder()

			aboutMeHandler(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("got status %d, want %d", w.Code, tt.expectedStatus)
			}

			if !strings.Contains(w.Body.String(), tt.expectedBody) {
				t.Errorf("got body %q, want to contain %q", w.Body.String(), tt.expectedBody)
			}
		})
	}
}
