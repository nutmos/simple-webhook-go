package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestWebhook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/webhook", strings.NewReader(`{"event":"ping"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), `"status":"received"`) {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}
