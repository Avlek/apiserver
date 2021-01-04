package impl

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s := NewAPIServer(NewConfig())
	s.router.GET("/hello", s.handleHello())
	s.router.ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
