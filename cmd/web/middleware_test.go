package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	secureHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()

	frameWant := "deny"
	frameOptions := rs.Header.Get("X-Frame-Options")
	if frameOptions != frameWant {
		t.Errorf("want %q; got %q", frameWant, frameOptions)
	}

	xssWant := "1; mode=block"
	xssProtection := rs.Header.Get("X-XSS-Protection")
	if xssProtection != xssWant {
		t.Errorf("want %q; got %q", xssWant, xssProtection)
	}

	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}
