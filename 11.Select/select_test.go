package selecting

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("tests that one url obtain should be faster than another", func(t *testing.T){
		slowServer := makeDelayedServer(30 * time.Millisecond)
		fastServer := makeDelayedServer(3 * time.Millisecond)

		defer 	slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL 

		want:= fastURL
		got := RacerSelect(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("tests that a timeout should be caught", func(t *testing.T){
		serverA:= makeDelayedServer(11 * time.Second)
		serverB:= makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil{
			t.Error("expeted an error but din't get it")
		}
	})
}

func makeDelayedServer(d time.Duration) *httptest.Server{
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(d)
				w.WriteHeader(http.StatusOK)
	}))
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

