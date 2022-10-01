package cmd

import (
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
