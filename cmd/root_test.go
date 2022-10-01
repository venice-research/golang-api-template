package cmd

import (
	"net/http"
	"testing"
	"time"
)

func TestStartupMessage(t *testing.T) {
	currentTime := time.Now()

	want := "GoLang API Template started at: " + currentTime.Format("01-02-2006")
	got := generateStartupMessage(currentTime)

	if want != got {
		t.Fatalf(`generateStartupMessage(currentTime) = %q, want match for %#q, nil`, got, want)
	}
}

func TestLogRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	want := "GET /"
	got := logRequest(req)

	if want != got {
		t.Fatalf(`logRequest(request) = %q, want match for %#q, nil`, got, want)
	}
}
