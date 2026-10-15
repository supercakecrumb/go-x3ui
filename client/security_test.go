package client

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestResponseBodiesDoNotLeak(t *testing.T) {
	const sentinel = "private-response-sentinel"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/login" {
			w.WriteHeader(500)
			io.WriteString(w, sentinel)
			return
		}
		io.WriteString(w, `{"success":false,"msg":"`+sentinel+`"}`)
	}))
	defer server.Close()
	var output bytes.Buffer
	c := NewClient(server.URL, "private-user-sentinel", "private-password-sentinel", false, slog.New(slog.NewTextHandler(&output, &slog.HandlerOptions{Level: slog.LevelDebug})))
	if err := c.Login(); err == nil || strings.Contains(err.Error(), sentinel) {
		t.Fatal("response body leaked from login error")
	}
	c.LastLoginTime = time.Now()
	if _, err := c.ListInbounds(); err == nil || strings.Contains(err.Error(), sentinel) {
		t.Fatal("response body leaked from list error")
	}
	if _, err := c.GetNewCertificate(); err == nil || strings.Contains(err.Error(), sentinel) {
		t.Fatal("response body leaked from certificate error")
	}
	for _, value := range []string{sentinel, "private-user-sentinel", "private-password-sentinel"} {
		if strings.Contains(output.String(), value) {
			t.Fatal("sensitive data leaked into logs")
		}
	}
}
