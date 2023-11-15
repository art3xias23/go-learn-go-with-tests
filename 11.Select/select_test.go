package selecting

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)
func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defer 	slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL 

	want:= fastURL
	got := RacerSelect(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

	func BenchmarkRacerSelect(b *testing.B){
		for i:=0; i<b.N; i++{
			RacerSelect("google.com", "amazon.com")
		}
	}

	func BenchmarkRacerTime(b *testing.B){
		for i:=0; i<b.N; i++{
			RacerTime("google.com", "amazon.com")
		}
	}

